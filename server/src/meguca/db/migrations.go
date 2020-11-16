package db

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"meguca/auth"
	"meguca/common"
	"meguca/config"
	"meguca/util"
	"strconv"
	"time"

	"github.com/go-playground/log"
	"github.com/lib/pq"
)

var version = len(migrations)

var migrations = []func(*sql.Tx) error{
	func(tx *sql.Tx) (err error) {
		// Initialize DB
		err = execAll(tx,
			`create table main (
				id text primary key,
				val text not null
			)`,
			`insert into main (id, val) values
				('version', '0'),
				('pyu', '0')`,
			`create table accounts (
				id varchar(20) primary key,
				password bytea not null
			)`,
			`create table sessions (
				account varchar(20) not null references accounts on delete cascade,
				token text not null,
				expires timestamp not null,
				primary key (account, token)
			)`,
			`create table bans (
				board varchar(3) not null,
				ip inet not null,
				by varchar(20) not null,
				reason text not null,
				expires timestamp default now(),
				primary key (ip, board)
			)`,
			`create table images (
				apng boolean not null,
				audio boolean not null,
				video boolean not null,
				fileType smallint not null,
				thumbType smallint not null,
				dims smallint[4] not null,
				length int not null,
				size int not null,
				md5 char(22) not null,
				sha1 char(40) primary key
			)`,
			`create table image_tokens (
				token char(86) not null primary key,
				sha1 char(40) not null references images on delete cascade,
				expires timestamp not null
			)`,
			`create table boards (
				readOnly boolean not null,
				textOnly boolean not null,
				forcedAnon boolean not null,
				hashCommands boolean not null,
				codeTags boolean not null,
				id varchar(3) primary key,
				ctr bigint default 0,
				created timestamp not null,
				title varchar(100) not null,
				notice varchar(500) not null,
				rules varchar(5000) not null,
				eightball text[] not null
			)`,
			`create table staff (
				board varchar(3) not null references boards on delete cascade,
				account varchar(20) not null references accounts on delete cascade,
				position varchar(50) not null
			)`,
			`create index staff_board on staff (board)`,
			`create index staff_account on staff (account)`,
			`create sequence post_id`,
			`create table threads (
				locked boolean,
				board varchar(3) not null references boards on delete cascade,
				id bigint primary key,
				postCtr bigint not null,
				imageCtr bigint not null,
				bumpTime bigint not null,
				replyTime bigint not null,
				subject varchar(100) not null,
				log text[] not null
			)`,
			`create index threads_board on threads (board)`,
			`create index bumpTime on threads (bumpTime)`,
			`create table posts (
				editing boolean not null,
				spoiler boolean,
				deleted boolean,
				banned boolean,
				id bigint primary key,
				op bigint not null references threads on delete cascade,
				time bigint not null,
				board varchar(3) not null,
				trip char(10),
				auth varchar(20),
				sha1 char(40) references images on delete set null,
				name varchar(50),
				imageName varchar(200),
				body varchar(2000) not null,
				password bytea,
				ip inet,
				links bigint[][2],
				backlinks bigint[][2],
				commands json[]
			)`,
			`create index deleted on posts (deleted)`,
			`create index op on posts (op)`,
			`create index image on posts (sha1)`,
			`create index editing on posts (editing)`,
			`create index ip on posts (ip)`,
		)
		if err != nil {
			return
		}

		data, err := json.Marshal(config.Defaults)
		if err != nil {
			return
		}
		q := sq.Insert("main").
			Columns("id", "val").
			Values("config", string(data))
		err = withTransaction(tx, q).Exec()
		if err != nil {
			return
		}

		err = CreateAdminAccount(tx)
		if err != nil {
			return
		}
		return CreateSystemAccount(tx)
	},
	func(tx *sql.Tx) (err error) {
		// Delete legacy columns
		return execAll(tx,
			`ALTER TABLE threads
				DROP COLUMN locked`,
			`ALTER TABLE boards
				DROP COLUMN hashCommands,
				DROP COLUMN codeTags`,
		)
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`ALTER TABLE threads
				DROP COLUMN log`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`ALTER TABLE boards
				DROP COLUMN ctr`,
		)
		return
	},
	// Restore correct image counters after incorrect updates
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`UPDATE threads
				SET imageCtr = (SELECT COUNT(*) FROM posts
					WHERE SHA1 IS NOT NULL
						AND op = threads.id
				)`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`ALTER TABLE images
				ADD COLUMN Title varchar(100) not null default '',
				ADD COLUMN Artist varchar(100) not null default ''`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`ALTER TABLE posts
				ADD COLUMN sage bool`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(`DROP INDEX deleted`)
		return
	},
	// Set default expiry configs, to keep all threads from deleting
	func(tx *sql.Tx) (err error) {
		return patchConfigs(tx, func(conf *config.Configs) {
			conf.ThreadExpiryMin = config.Defaults.ThreadExpiryMin
			conf.ThreadExpiryMax = config.Defaults.ThreadExpiryMax
		})
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`ALTER TABLE boards
				ADD COLUMN disableRobots bool default false`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`ALTER TABLE threads
				ADD COLUMN sticky bool default false`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`ALTER TABLE bans
				ADD COLUMN forPost bigint default 0`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		return execAll(tx,
			`create table mod_log (
				type smallint not null,
				board varchar(3) not null,
				id bigint not null,
				by varchar(20) not null,
				created timestamp default (now() at time zone 'utc')
			)`,
			`create index mod_log_board on mod_log (board)`,
			`create index mod_log_created on mod_log (created)`,
		)
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(`create index sticky on threads (sticky)`)
		return
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`ALTER TABLE posts
				DROP COLUMN backlinks`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`create table banners (
				board varchar(3) not null references boards on delete cascade,
				id smallint not null,
				data bytea not null,
				mime text not null
			);`,
		)
		return
	},
	func(tx *sql.Tx) error {
		return execAll(tx,
			`alter table boards
				alter column id type text`,
			`alter table bans
				alter column board type text`,
			`alter table mod_log
				alter column board type text`,
			`alter table staff
				alter column board type text`,
			`alter table banners
				alter column board type text`,
			`alter table threads
				alter column board type text`,
			`alter table posts
				alter column board type text`,
		)
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`create table loading_animations (
				board text primary key references boards on delete cascade,
				data bytea not null,
				mime text not null
			);`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`ALTER TABLE boards
				ADD COLUMN defaultCSS text default 'moe'`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`ALTER TABLE posts
				ADD COLUMN flag char(2)`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`ALTER TABLE boards
				ADD COLUMN flags bool default false`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`alter table images
				alter column title type varchar(200)`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`ALTER TABLE boards
				ADD COLUMN NSFW bool default false`,
		)
		return
	},
	func(tx *sql.Tx) error {
		return execAll(tx,
			`create table reports (
				id bigserial primary key,
				target bigint not null,
				board text not null,
				reason text not null,
				by inet not null,
				illegal boolean not null,
				created timestamp default (now() at time zone 'utc')
			)`,
			`create index report_board on reports (board)`,
			`create index report_created on reports (created)`,
		)
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`ALTER TABLE boards
				ADD COLUMN nonLive bool default false`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`ALTER TABLE threads
				ADD COLUMN nonLive bool default false`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`ALTER TABLE boards
				ADD COLUMN posterIDs bool default false`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`ALTER TABLE posts
				ADD COLUMN posterID text`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`ALTER TABLE threads
				ADD COLUMN locked bool default false`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`ALTER TABLE boards
				ADD COLUMN js varchar(5000) default ''`,
		)
		return
	},
	// Fix consequences of bug in init.sql
	func(tx *sql.Tx) (err error) {
		q := `SELECT EXISTS (SELECT 1
			FROM information_schema.columns
			WHERE table_schema='public'
				AND table_name='posts'
				AND column_name='locked'
		)`
		var exists bool
		err = tx.QueryRow(q).Scan(&exists)
		if err != nil || !exists {
			return
		}

		// Correct it
		_, err = tx.Exec(
			`ALTER TABLE posts
				DROP COLUMN locked`,
		)
		if err != nil {
			return
		}
		_, err = tx.Exec(
			`ALTER TABLE threads
				ADD COLUMN locked bool default false`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`ALTER TABLE boards
				DROP COLUMN js`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		return execAll(tx,
			`DROP INDEX editing`,
			`ALTER TABLE boards
				DROP COLUMN nonLive`,
			`ALTER TABLE posts
				DROP COLUMN editing`,
		)
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`ALTER TABLE posts
				DROP COLUMN password`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		return execAll(tx,
			`ALTER TABLE boards
				ADD COLUMN nonLive bool default false`,
			`ALTER TABLE posts
				ADD COLUMN 	editing boolean not null default false`,
			`create index editing on posts (editing);`,
		)
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`ALTER TABLE posts
				ADD COLUMN password bytea`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`INSERT INTO main VALUES ('roulette', '6')`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		// Fuck any wise guy trying to create an account nad block an upgrade
		_, err = tx.Exec(
			`DELETE FROM accounts
			WHERE id = 'system'`,
		)
		if err != nil {
			return
		}

		password, err := auth.RandomID(32)
		if err != nil {
			return
		}
		hash, err := auth.BcryptHash(password, 10)
		if err != nil {
			return
		}
		_, err = tx.Exec(
			`insert into accounts (id, password)
			values ('system', $1)`,
			hash,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		return patchConfigs(tx, func(conf *config.Configs) {
			conf.CharScore = config.Defaults.CharScore
			conf.PostCreationScore = config.Defaults.PostCreationScore
			conf.ImageScore = config.Defaults.ImageScore
		})
	},
	func(tx *sql.Tx) (err error) {
		var rcount string
		err = tx.QueryRow(
			`SELECT COUNT(*) FROM mod_log WHERE board != 'all' AND by = 'system' AND type = 0`,
		).Scan(&rcount)
		if err != nil {
			return
		}
		_, err = tx.Exec(
			`INSERT INTO main VALUES ('rcount', $1)`,
			rcount,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		execAll(tx,
			`create table links (
				source bigint not null references posts on delete cascade,
				target bigint not null references posts on delete cascade,
				primary key(source, target)
			);`,
			`create index link_source on links (source);`,
			`create index link_target on links (target);`,
		)
		if err != nil {
			return
		}

		// Read all posts and links
		r, err := tx.Query(
			`select p.id, p.op, p.links from posts as p
			join threads as t on t.id = p.op`,
		)
		if err != nil {
			return
		}
		defer r.Close()
		var (
			posts  = make(map[uint64]bool, 1<<10)
			links  = make(map[uint64]uint64, 1<<10)
			id, op uint64
			lr     linkRowLegacy
		)
		for r.Next() {
			err = r.Scan(&id, &op, &lr)
			if err != nil {
				return
			}
			posts[id] = true
			for _, pair := range lr {
				links[id] = pair[0]
			}
		}
		if err != nil {
			return
		}

		// Remove legacy link row
		_, err = tx.Exec(`alter table posts drop column links`)
		if err != nil {
			return
		}

		// Write only verified links to new table
		q, err := tx.Prepare(
			`insert into links (source, target)
			values ($1, $2)`,
		)
		if err != nil {
			return
		}
		for source, target := range links {
			if !posts[source] || !posts[target] {
				continue
			}
			_, err = q.Exec(source, target)
			if err != nil {
				return
			}
		}

		return
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(`ALTER TABLE banners DROP COLUMN id`)
		return
	},
	func(tx *sql.Tx) error {
		return execAll(tx,
			`DELETE FROM mod_log`,
			`DELETE FROM bans`,
			`ALTER TABLE mod_log ADD CONSTRAINT mod_log_board_fkey
			FOREIGN KEY (board) REFERENCES boards(id) ON DELETE CASCADE`,
			`ALTER TABLE bans ADD CONSTRAINT bans_board_fkey
			FOREIGN KEY (board) REFERENCES boards(id) ON DELETE CASCADE`,
		)
	},
	func(tx *sql.Tx) error {
		return execAll(tx,
			`ALTER TABLE threads DROP COLUMN postCtr`,
			`ALTER TABLE threads DROP COLUMN imageCtr`,
		)
	},
	func(tx *sql.Tx) error {
		return execAll(tx,
			`ALTER TABLE mod_log ADD COLUMN length bigint default 0`,
			`ALTER TABLE mod_log ADD COLUMN reason text default ''`,
		)
	},
	func(tx *sql.Tx) error {
		return execAll(tx,
			`create index image_fileType on images (fileType)`,
			`create index image_audio on images (audio)`,
			`create index post_board on posts (board)`,
		)
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`ALTER TABLE boards
				ADD COLUMN rbText bool default false`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		// Read all commands
		var (
			comms = make(map[uint64][]common.Command, 1024)
			id    uint64
		)
		err = queryAll(
			withTransaction(tx,
				sq.Select("id", "commands").
					From("posts").
					Where("commands is not null")),
			func(r *sql.Rows) (err error) {
				var com commandRow
				err = r.Scan(&id, &com)
				if err != nil {
					return
				}
				comms[id] = []common.Command(com)
				return
			},
		)
		if err != nil {
			return
		}

		prep, err := tx.Prepare(
			`update posts
			set commands = $2
			where id = $1`)
		if err != nil {
			return
		}

		// Remove all #pyu/#pcount commands
		new := make(commandRow, 0, 64)
		for id, comms := range comms {
			new = new[:0]
			for _, c := range comms {
				switch c.Type {
				case common.Pyu, common.Pcount:
				default:
					new = append(new, c)
				}
			}
			val := new
			if len(new) == 0 {
				val = nil
			}
			_, err = prep.Exec(id, val)
			if err != nil {
				return
			}
		}

		return
	},
	func(tx *sql.Tx) error {
		return withTransaction(tx, sq.Delete("main").Where("id = 'pyu'")).Exec()
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`ALTER TABLE boards
				ADD COLUMN pyu bool default false`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`create table pyu (
				id text primary key references boards on delete cascade,
				pcount bigint default 0
			);
			create table pyu_limit (
				ip inet not null,
				board text not null references boards on delete cascade,
				expires timestamp not null,
				pcount smallint default 4,
				primary key(ip, board)
			);
			create index pyu_limit_ip on pyu_limit (ip);
			create index pyu_limit_board on pyu_limit (board);`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		// Reverted migration
		return
	},
	func(tx *sql.Tx) (err error) {
		// Revert changes
		return execAll(tx,
			`alter table boards drop column pyu`,
			`drop table pyu`,
			`drop table pyu_limit`,
		)
	},
	func(tx *sql.Tx) (err error) {
		return patchConfigs(tx, func(conf *config.Configs) {
			conf.EmailErrMail = config.Defaults.EmailErrMail
			conf.EmailErrPass = config.Defaults.EmailErrPass
			conf.EmailErrSub = config.Defaults.EmailErrSub
			conf.EmailErrPort = config.Defaults.EmailErrPort
		})
	},
	// Fixes global moderation
	func(tx *sql.Tx) (err error) {
		c := BoardConfigs{
			BoardConfigs: config.AllBoardConfigs.BoardConfigs,
			Created:      time.Now().UTC(),
		}

		err = withTransaction(tx, sq.Insert("boards").
			Columns(
				"id", "readOnly", "textOnly", "forcedAnon", "disableRobots",
				"flags", "NSFW",
				"posterIDs", "rbText", "created", "defaultCSS", "title", "notice",
				"rules", "eightball").
			Values(
				c.ID, c.ReadOnly, c.TextOnly, c.ForcedAnon, c.DisableRobots,
				c.Flags, c.NSFW, c.PosterIDs, c.RbText,
				c.Created, c.DefaultCSS, c.Title, c.Notice, c.Rules,
				pq.StringArray(c.Eightball))).
			Exec()

		if err != nil {
			return
		}

		err = notifyBoardUpdated(tx, c.ID)

		if err != nil {
			return
		}

		return WriteStaff(tx, "all", map[string][]string{
			"owners": {"admin", "system"},
		})
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`ALTER TABLE boards
				ADD COLUMN pyu bool default false`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`create table pyu (
				id text primary key references boards on delete cascade,
				pcount bigint default 0
			);
			create table pyu_limit (
				ip inet not null,
				board text not null references boards on delete cascade,
				expires timestamp not null,
				pcount smallint default 4,
				primary key(ip, board)
			);
			create index pyu_limit_ip on pyu_limit (ip);
			create index pyu_limit_board on pyu_limit (board);`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		r, err := tx.Query(`select id from boards`)
		if err != nil {
			return
		}
		defer r.Close()

		var boards []string
		for r.Next() {
			var board string
			err = r.Scan(&board)
			if err != nil {
				return
			}
			boards = append(boards, board)
		}
		err = r.Err()
		if err != nil {
			return
		}

		q, err := tx.Prepare(`insert into pyu (id, pcount) values ($1, 0)`)
		if err != nil {
			return
		}
		for _, b := range boards {
			_, err = q.Exec(b)
			if err != nil {
				return
			}
		}

		return
	},
	func(tx *sql.Tx) (err error) {
		return execAll(tx,
			`alter table pyu_limit drop column expires`,
			`alter table pyu_limit add column restricted bool default false`,
		)
	},
	func(tx *sql.Tx) (err error) {
		return execAll(tx,
			`create table captchas (
				id text primary key not null,
				solution text not null,
				expires timestamp not null
			)`,
			`create table failed_captchas (
				ip inet not null,
				expires timestamp not null
			)`,
		)
	},
	func(tx *sql.Tx) (err error) {
		return execAll(tx,
			`create index failed_captchas_ip on failed_captchas (ip)`,
		)
	},
	func(tx *sql.Tx) (err error) {
		var tasks []string
		for _, t := range [...]string{
			"image_tokens", "bans", "captchas", "failed_captchas",
		} {
			tasks = append(tasks, createIndex(t, "expires"))
		}
		tasks = append(tasks, createIndex("posts", "time"))
		return execAll(tx, tasks...)
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`INSERT INTO main VALUES ('geo_md5', 'initial value, ignore')`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		err = withTransaction(tx, sq.Delete("main").Where("id = 'roulette'")).Exec()

		if err != nil {
			return
		}

		return withTransaction(tx, sq.Delete("main").Where("id = 'rcount'")).Exec()
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`create table roulette (
				id bigint primary key references threads on delete cascade,
				scount smallint default 6,
				rcount smallint default 0
			);
			create index roulette_rcount on roulette (rcount);`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		var threads []uint64
		r, err := tx.Query(`select id from threads`)

		if err != nil {
			return
		}

		defer r.Close()

		for r.Next() {
			var thread uint64
			err = r.Scan(&thread)

			if err != nil {
				return
			}

			threads = append(threads, thread)
		}

		err = r.Err()

		if err != nil {
			return
		}

		q, err := tx.Prepare(
			`insert into roulette (id, scount, rcount) values ($1, 6, 0)`)

		if err != nil {
			return
		}

		for _, t := range threads {
			_, err = q.Exec(t)

			if err != nil {
				return
			}
		}

		return
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`create table spam_scores (
				ip inet primary key,
				score bigint not null
			);`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		return execAll(tx,
			`alter table boards drop column nonLive`,
			`alter table threads drop column nonLive`,
		)
	},
	func(tx *sql.Tx) (err error) {
		_, err = tx.Exec(
			`alter table posts add column meidoVision bool default false`,
		)
		return
	},
	func(tx *sql.Tx) (err error) {
		var tasks []string

		for _, col := range [...]string{"deleted", "banned", "meidovision"} {
			tasks = append(tasks, "alter table posts drop column "+col)
		}
		tasks = append(tasks,
			`alter table posts add column moderated bool not null default false`,
			`create table post_moderation (
				post_id bigint not null references posts on delete cascade,
				type smallint not null,
				by text not null,
				length bigint not null,
				reason text not null
			)`,
			createIndex("post_moderation", "post_id"),
		)

		return execAll(tx, tasks...)
	},
	func(tx *sql.Tx) (err error) {
		var tasks []string
		setNotNull := func(col, typ, def string) {
			tasks = append(tasks, fmt.Sprintf(
				`UPDATE posts
				SET %s = %s
				WHERE %s IS NULL`,
				col, def, col,
			))
			for _, s := range [...]string{
				"SET DATA TYPE " + typ,
				"SET DEFAULT " + def,
				"SET NOT NULL",
			} {
				tasks = append(tasks, fmt.Sprintf(
					`ALTER TABLE posts ALTER COLUMN %s %s`, col, s))
			}
		}

		for _, col := range [...]string{"spoiler", "sage"} {
			setNotNull(col, "bool", "false")
		}
		for _, col := range [...]string{
			"name", "trip", "auth", "imageName", "flag", "posterID",
		} {
			setNotNull(col, "text", `''::text`)
		}

		return execAll(tx, tasks...)
	},
	func(tx *sql.Tx) (err error) {
		for _, t := range [...]string{"mod_log", "post_moderation"} {
			_, err = tx.Exec(
				fmt.Sprintf(`alter table %s rename column reason to data`, t))
			if err != nil {
				return
			}
		}
		return
	},
	func(tx *sql.Tx) (err error) {
		return execAll(tx,
			`alter table bans drop constraint bans_pkey`,
			createIndex("bans", "ip"),
			createIndex("bans", "board"),
		)
	},
	func(tx *sql.Tx) (err error) {
		err = patchConfigs(tx, func(conf *config.Configs) {
			conf.CaptchaTags = config.Defaults.CaptchaTags
			conf.OverrideCaptchaTags = map[string]string{}
		})
		if err != nil {
			return
		}
		_, err = tx.Exec(`drop table captchas`)
		return
	},
	func(tx *sql.Tx) error {
		return execAll(tx,
			`create table last_solved_captchas (
				ip inet primary key,
				time timestamp not null default (now() at time zone 'utc')
			)`,
			createIndex("last_solved_captchas", "time"),
		)
	},
	func(tx *sql.Tx) error {
		return execAll(tx,
			loadSQL("triggers/notify_thread_post_count"),
			`create trigger notify_thread_post_count
			after insert
			on posts
			for each row
			execute procedure notify_thread_post_count()`,

			loadSQL("triggers/notify_thread_deleted"),
			`create trigger notify_thread_deleted
			after delete
			on threads
			for each row
			execute procedure notify_thread_deleted()`,
		)
	},
	func(tx *sql.Tx) error {
		// Apply function changes
		return execAll(tx,
			loadSQL("triggers/notify_thread_post_count"),
			loadSQL("triggers/notify_thread_deleted"),
		)
	},
}

func createIndex(table, column string) string {
	return fmt.Sprintf(`create index %s_%s on %s (%s)`, table, column, table,
		column)
}

func loadSQL(path string) string {
	return string(MustAsset(path + ".sql"))
}

// Run migrations from version `from`to version `to`
func runMigrations(from, to int) (err error) {
	var tx *sql.Tx
	for i := from; i < to; i++ {
		if !IsTest {
			log.Infof("upgrading database to version %d", i+1)
		}
		tx, err = db.Begin()
		if err != nil {
			return
		}

		err = migrations[i](tx)
		if err != nil {
			return rollBack(tx, err)
		}

		// Write new version number
		_, err = tx.Exec(
			`update main set val = $1 where id = 'version'`,
			i+1,
		)
		if err != nil {
			return rollBack(tx, err)
		}

		err = tx.Commit()
		if err != nil {
			return
		}
	}
	return
}

func rollBack(tx *sql.Tx, err error) error {
	if rbErr := tx.Rollback(); rbErr != nil {
		err = util.WrapError(err.Error(), rbErr)
	}
	return err
}

// Patches server configuration during upgrades
func patchConfigs(tx *sql.Tx, fn func(*config.Configs)) (err error) {
	var s string
	err = tx.QueryRow("SELECT val FROM main WHERE id = 'config'").Scan(&s)
	if err != nil {
		return
	}
	conf, err := decodeConfigs(s)
	if err != nil {
		return
	}

	fn(&conf)

	buf, err := json.Marshal(conf)
	if err != nil {
		return
	}
	_, err = tx.Exec(
		`UPDATE main
			SET val = $1
			WHERE id = 'config'`,
		string(buf),
	)
	return
}

// For decoding and encoding the tuple arrays we used to store links in.
// Still needed for migrations.
type linkRowLegacy [][2]uint64

func (l *linkRowLegacy) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		return l.scanBytes(src)
	case string:
		return l.scanBytes([]byte(src))
	case nil:
		*l = nil
		return nil
	default:
		return fmt.Errorf("cannot convert %T to [][2]uint", src)
	}
}

func (l *linkRowLegacy) scanBytes(src []byte) error {
	length := len(src)
	if length < 6 {
		return errors.New("source too short")
	}

	src = src[1 : length-1]

	// Determine needed size and preallocate final array
	commas := 0
	for _, b := range src {
		if b == ',' {
			commas++
		}
	}
	*l = make(linkRowLegacy, 0, (commas-1)/2+1)

	var (
		inner bool
		next  [2]uint64
		err   error
		buf   = make([]byte, 0, 16)
	)
	for _, b := range src {
		switch b {
		case '{': // New tuple
			inner = true
			buf = buf[0:0]
		case ',':
			if inner { // End of first uint
				next[0], err = strconv.ParseUint(string(buf), 10, 64)
				if err != nil {
					return err
				}
				buf = buf[0:0]
			}
		case '}': // End of tuple
			next[1], err = strconv.ParseUint(string(buf), 10, 64)
			if err != nil {
				return err
			}
			*l = append(*l, next)
		default:
			buf = append(buf, b)
		}
	}

	return nil
}

func (l linkRowLegacy) Value() (driver.Value, error) {
	n := len(l)
	if n == 0 {
		return nil, nil
	}

	b := make([]byte, 1, 16)
	b[0] = '{'
	for i, l := range l {
		if i != 0 {
			b = append(b, ',')
		}
		b = append(b, '{')
		b = strconv.AppendUint(b, l[0], 10)
		b = append(b, ',')
		b = strconv.AppendUint(b, l[1], 10)
		b = append(b, '}')
	}
	b = append(b, '}')

	return string(b), nil
}
