# wik
###### *the transparent wiki*

So, [`gitit`](http://gitit.net/) still isn't compiling under `nix`. And I'm still bored. And I was curious about [these benchmark results](https://github.com/fukamachi/woo#how-fast). So here's a minimal wiki in [Go](http://golang.org/).

### Setup

- Install `git`
- Install the [`blackfriday` markdown parser](https://github.com/russross/blackfriday) and the [`bluemonday` html sanitizer](https://github.com/microcosm-cc/bluemonday)
- Clone this repo, and `cd` into it
- Initialize a wiki repo somewhere
- Run `go build && ./wik path/to/your/wiki/repo/here/`

### Usage

###### Creating a Page

- Visit the URI of a nonexistent page
- Click the `Create` button (this will create nested directories as appropriate)

###### Editing a Page

- Visit a page
- Click the `Edit` button
- This will show a page with a `codemirror` instance which you can use to edit the page
- Click `Save` when done (you can also discard changes by clicking the `Cancel` link)

###### Removing a Page

- Visit a page
- Click the `Remove` button (list pages are auto-generated and so can't be deleted)

### License & Credits

`wik` is released under the [AGPL3](http://www.gnu.org/licenses/agpl-3.0.html). It also includes some media under compatible licenses:

- Icons from [the Silk icon set](http://commons.wikimedia.org/wiki/Category:Silk_icons), released under CC-BY 2.5
- [CodeMirror](http://codemirror.net/) and the associated [`markdown`](http://daringfireball.net/projects/markdown/) highlighting mode. `markdown` is released under a [BSD-style license](http://daringfireball.net/projects/markdown/license), while CodeMirror and its `markdown` implementation are released under an [Expat-style license](http://codemirror.net/LICENSE)
