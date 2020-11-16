#include "mutations.hh"
#include <emscripten.h>
#include <optional>
#include <unordered_map>
#include <unordered_set>
#include <vector>

namespace brunhild {
using std::string;

// Pending mutations for an element
struct Mutations {
    bool remove_el = false, scroll_into_view = false;
    std::optional<std::string> set_inner_html, set_outer_html;
    std::vector<std::string> append, prepend, before, after, move_prepend,
        move_after;
    std::unordered_set<std::string> remove_attr;
    std::unordered_map<std::string, std::string> set_attr;

    // Clear mutations of element inner content to free up memory
    void free_inner();

    // Clear mutations of element inner and outer content to free up memory
    void free_outer();

    // Execute buffered mutations
    void exec(const std::string& id);
};

void (*before_flush)() = nullptr;
void (*after_flush)() = nullptr;

// TODO: Should probably use some smarter data structure, that maintains both
// mutation set insertion order and is searchable by string

// All pending mutations quickly accessible by element ID
static std::unordered_map<string, Mutations> mutations;

// Stores mutation order, so we can somewhat make sure, new children are not
// manipulated, before insertion
static std::vector<std::string> mutation_order;

// Fetches a mutation set by element ID or creates a new one ond registers
// its execution order
static Mutations* get_mutation_set(string id)
{
    if (!mutations.count(id)) {
        mutation_order.push_back(id);
    }
    return &mutations[id];
}

void append(string id, string html)
{
    get_mutation_set(id)->append.push_back(html);
}

void prepend(string id, string html)
{
    get_mutation_set(id)->prepend.push_back(html);
}

void before(string id, string html)
{
    get_mutation_set(id)->before.push_back(html);
}

void after(string id, string html)
{
    get_mutation_set(id)->after.push_back(html);
}

// Move child node to the front of the parent
void move_prepend(string parent_id, string child_id)
{
    get_mutation_set(parent_id)->move_prepend.push_back(child_id);
}

// Move child node after a sibling in the parent
void move_after(string sibling_id, string child_id)
{
    get_mutation_set(sibling_id)->move_prepend.push_back(child_id);
}

void set_inner_html(string id, string html)
{
    auto mut = get_mutation_set(id);
    // These would be overwritten, so we can free up used memory
    mut->free_inner();
    mut->set_inner_html = html;
}

void set_outer_html(string id, string html)
{
    auto mut = get_mutation_set(id);
    mut->free_outer();
    mut->set_outer_html = html;
}

void remove(string id)
{
    auto mut = get_mutation_set(id);
    mut->free_outer();
    mut->remove_el = true;
}

void set_attr(string id, string key, string val)
{
    get_mutation_set(id)->set_attr[key] = val;
}

void remove_attr(string id, string key)
{
    auto mut = get_mutation_set(id);
    mut->set_attr.erase(key);
    mut->remove_attr.insert(id);
}

void scroll_into_view(string id)
{
    get_mutation_set(id)->scroll_into_view = true;
}

void Mutations::free_inner()
{
    append.clear();
    prepend.clear();
    move_prepend.clear();
    move_after.clear();
    set_inner_html = std::nullopt;
}

void Mutations::free_outer()
{
    free_inner();
    remove_attr.clear();
    set_attr.clear();
    set_outer_html = std::nullopt;
}

extern "C" void flush()
{
    if (before_flush) {
        (*before_flush)();
    }

    if (mutations.size()) {
        for (auto& id : mutation_order) {
            mutations.at(id).exec(id);
        }
        mutation_order.clear();
        mutations.clear();
    }

    if (after_flush) {
        (*after_flush)();
    }
}

void Mutations::exec(const string& id)
{
    // Assign element to global variable, so we don't have to look it up each
    // time
    const bool exists = (bool)EM_ASM_INT(
        {
            window.__el = document.getElementById(UTF8ToString($0));
            return !!window.__el;
        },
        id.c_str());
    if (!exists) {
        // Nothing we can do
        return;
    }

    // TODO: Do these loops in one JS call, if possible

    // Before and after inserts need to happen, even if the element is going to
    // be removed
    for (auto& html : before) {
        EM_ASM_INT(
            {
                var el = window.__el;
                var cont = document.createElement('div');
                cont.innerHTML = UTF8ToString($0);
                el.parentNode.insertBefore(cont.firstChild, el);
            },
            html.c_str());
    }
    for (auto& html : after) {
        EM_ASM_INT(
            {
                var el = window.__el;
                var cont = document.createElement('div');
                cont.innerHTML = UTF8ToString($0);
                el.parentNode.insertBefore(cont.firstChild, el.nextSibling);
            },
            html.c_str());
    }

    if (remove_el) {
        EM_ASM({
            var el = window.__el;
            el.parentNode.removeChild(el);
        });
        // If the element is to be removed, nothing else needs to be done
        return;
    }

    if (set_outer_html) {
        EM_ASM_INT({ window.__el.outerHTML = UTF8ToString($0); },
            set_outer_html->c_str());
    }
    if (set_inner_html) {
        EM_ASM_INT({ window.__el.innerHTML = UTF8ToString($0); },
            set_inner_html->c_str());
    }

    for (auto& html : append) {
        EM_ASM_INT(
            {
                var el = window.__el;
                var cont = document.createElement('div');
                cont.innerHTML = UTF8ToString($0);
                el.appendChild(cont.firstChild);
            },
            html.c_str());
    }
    for (auto& html : prepend) {
        EM_ASM_INT(
            {
                var el = window.__el;
                var cont = document.createElement('div');
                cont.innerHTML = UTF8ToString($0);
                el.insertBefore(cont.firstChild, el.firstChild);
            },
            html.c_str());
    }
    for (auto& child_id : move_prepend) {
        EM_ASM_INT(
            {
                var el = window.__el;
                el.insertBefore(
                    document.getElementById(UTF8ToString($0)), el.firstChild);
            },
            child_id.c_str());
    }
    for (auto& child_id : move_after) {
        EM_ASM_INT(
            {
                var el = window.__el;
                el.parentNode.insertBefore(
                    document.getElementById(UTF8ToString($0)), el.nextSibling);
            },
            child_id.c_str());
    }

    for (auto& kv : set_attr) {
        EM_ASM_INT(
            { window.__el.setAttribute(UTF8ToString($0), UTF8ToString($1)); },
            kv.first.c_str(), kv.second.c_str());
    }
    for (auto& key : remove_attr) {
        EM_ASM_INT(
            { window.__el.removeAttribute(UTF8ToString($0)); }, key.c_str());
    }

    if (scroll_into_view) {
        EM_ASM({ window.__el.scrollIntoView(); });
    }
}
}
