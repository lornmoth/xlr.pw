#include "etc.hh"
#include "../lang.hh"
#include "../options/options.hh"
#include "../state.hh"
#include "../util.hh"
#include "view.hh"
#include <sstream>

using std::string;
using std::string_view;

// Renders "56 minutes ago" or "in 56 minutes" like relative time text
// Units is the index used to retrieve the language pack value for unit
// pluralization.
static string ago(time_t n, string units, bool is_future)
{
    auto count = pluralize(n, units);
    return is_future ? lang.posts.at("in") + " " + count
                     : count + " " + lang.posts.at("ago");
}

string relative_time(time_t then)
{
    auto now = (float)std::time(0);
    auto t = (now - (float)then) / 60;
    auto is_future = false;
    if (t < 1) {
        if (t > -5) { // Assume to be client clock imprecision
            return lang.posts.at("justNow");
        }
        is_future = true;
        t = -t;
    }

    const int divide[4] = { 60, 24, 30, 12 };
    const static string unit[4] = { "minute", "hour", "day", "month" };
    for (int i = 0; i < 4; i++) {
        if (t < divide[i]) {
            return ago(t, unit[i], is_future);
        }
        t /= divide[i];
    }

    return ago(t, "year", is_future);
}

std::string absolute_thread_url(unsigned long id, string board)
{
    std::ostringstream s;
    s << '/' << board << '/' << id;
    return s.str();
}

Node render_post_link(unsigned long id, const LinkData& data)
{
    const bool cross_thread = data.op != page.thread;
    const string id_str = std::to_string(id);

    std::ostringstream url;
    if (cross_thread) {
        url << '/' << data.board << '/' << data.op;
    }
    url << "#p" << id_str;

    std::ostringstream text;
    text << ">>" << id_str;
    if (cross_thread && page.thread) {
        text << " ➡";
    }
    if (post_ids.mine.count(id)) { // Post, the user made
        text << ' ' << lang.posts.at("you");
    }

    Node n = Node("em");
    n.children.reserve(2);
    string cls = "post-link";
    string hash_cls = "hash-link";
    if (post_ids.hidden.count(id)) {
        cls += " strikethrough";
        hash_cls += " strikethrough";
    }
    n.children.push_back({ "a",
        {
            { "class", cls }, { "href", url.str() },
        },
        text.str() });
    if (options.post_inline_expand) {
        n.children.push_back({ "a",
            {
                { "class", hash_cls }, { "href", url.str() },
            },
            " #" });
    }

    return n;
}

Node render_link(string_view url, string_view text, bool new_tab)
{
    Node n({
        "a",
        {
            { "rel", "noreferrer" }, { "href", brunhild::escape(string(url)) },
        },
        string(text), true,
    });
    if (new_tab) {
        n.attrs["target"] = "_blank";
    }
    return n;
}

Post* match_post(emscripten::val& event)
{
    string attr
        = event["target"].call<string>("getAttribute", string("data-id"));
    if (attr == "") {
        return 0;
    }
    const unsigned long id = std::stoul(attr);
    if (!posts.count(id)) {
        return 0;
    }
    return &posts.at(id);
}

std::optional<std::tuple<Post*, PostView*>> match_view(emscripten::val& event)
{
    auto model = match_post(event);
    if (!model) {
        return {};
    }
    const string id = event["target"]
                          .call<emscripten::val>("closest", string("article"))
                          .call<string>("getAttribute", string("id"));
    for (auto& v : model->views) {
        if (v->id == id) {
            return { { model, v.get() } };
        }
    }
    return {};
}
