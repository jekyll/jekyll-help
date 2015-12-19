# [Jekyll/GitHub workflow with hidden drafts?](https://github.com/jekyll/jekyll-help/issues/27)

> state: **closed** opened by: **** on: ****

I am thinking about doing a blog/site with multiple authors, and while I would like for the site to be open source (at least the majority of it), I would prefer the drafts remained hidden, and that, when published as posts, their edit history would be erased from before they were published.

&#x60;&#x60;&#x60;
private repo:
  - _drafts/ (w/ edit history)

public repo:
  - _posts/ (w/ edit history)
  - (...)
&#x60;&#x60;&#x60;

Assuming everything is hosted on GitHub, but deployed to S3 (to support custom domains), what would be the best way to do this?

And which publication hook would fit the best for automating this process as much as possible?

### Comments

---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/27#issuecomment-42128583) on: **Sunday, May 04, 2014**

gitignore the drafts folder? Will more than one person be working on the
same draft? Also, for what it&#x27;s worth, GitHub Pages supports custom domains
via a CNAME file in the repo. See
https://help.github.com/articles/setting-up-a-custom-domain-with-github-pages


On 4 May 2014 18:56, ndarville &lt;notifications@github.com&gt; wrote:

&gt; I am thinking about doing a blog/site with multiple authors, and while I
&gt; would like for the site to be open source (at least the majority of it), I
&gt; would prefer the drafts remained hidden, and that, when published as posts,
&gt; their edit history would be erased from before they were published.
&gt;
&gt; private repo:
&gt;   - _drafts/
&gt;
&gt; public repo:
&gt;   - _posts
&gt;   - (...)
&gt;
&gt; Assuming everything is hosted on GitHub, but deployed to S3 (to support
&gt; custom domains), what would be the best way to do this?
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub&lt;https://github.com/jekyll/jekyll-help/issues/27&gt;
&gt; .
&gt;
---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/27#issuecomment-42128632) on: **Sunday, May 04, 2014**

One person will mainly focus on the draft, but I still want it to be a collaborative process where people comment using the Issues system and helps out in other ways—not to mention if the process involves editors, graphic department, fact-checking, proofing, and so on.

So syncing the draft centrally is still important.
---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/27#issuecomment-42128683) on: **Sunday, May 04, 2014**

In that case, if you&#x27;re using git, it will always see a removal from
_drafts and inclusion in _posts in the same commit as a &#x27;move&#x27; operation as
git doesn&#x27;t actually have a move/copy command it uses a heuristic to detect
when a file is copied/moved in the same commit based on the contents of the
files. Your best bet is probably to have one commit remove the file in
_drafts and another add it back into _posts if you *really* want to make
sure there&#x27;s no history attached. My question would be why you want this
though.


On 4 May 2014 19:34, ndarville &lt;notifications@github.com&gt; wrote:

&gt; One person will mainly focus on the draft, but I still want it to be a
&gt; collaborative process where people comment using the Issues system and
&gt; helps out in other ways—not to mention if the process involves editors,
&gt; fact-checking, and proofing. So syncing the draft is still important.
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub&lt;https://github.com/jekyll/jekyll-help/issues/27#issuecomment-42128632&gt;
&gt; .
&gt;
---
> from: [**mhulse**](https://github.com/jekyll/jekyll-help/issues/27#issuecomment-42135446) on: **Sunday, May 04, 2014**

I wonder if you could put the drafts in a private repo (I think free accounts are allowed one private free of charge) and the use a symlink to pull drafts into the jekyll repo whilst using a&#x60;.gitignore&#x60; to keep &#x60;_drafts&#x60; from showing in commit history?

A more dumbed down approach could be something simular to above but use Dropbox to share the drafts with collaborators?
---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/27#issuecomment-42135524) on: **Sunday, May 04, 2014**

Yeah, if the &#x60;.gitignore&#x60; thing works with importing the private repo as a submodule or some such, that’d probably work pretty well.
---
> from: [**mhulse**](https://github.com/jekyll/jekyll-help/issues/27#issuecomment-42135755) on: **Sunday, May 04, 2014**

You could maybe use a git submodule but I&#x27;ve found that symlinks are pretty handy in terms of instantaneous updates. Not sure, though, how Jekyll will handle symlink to outside repo/files?
---
> from: [**mhulse**](https://github.com/jekyll/jekyll-help/issues/27#issuecomment-42135799) on: **Sunday, May 04, 2014**

On the other hand, if you don&#x27;t need to **instantaneously** preview the drafts on the local site, a submodule might be best (like you&#x27;re already saying). :smile:
---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/27#issuecomment-42136504) on: **Sunday, May 04, 2014**

True, it does mess with the ease of viewing the editing history. A submodule probably isn’t the way to go.
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/27#issuecomment-42277116) on: **Tuesday, May 06, 2014**

@ndarville Read this one https://github.com/jekyll/jekyll/issues/1469#issuecomment-23831358
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/27#issuecomment-55624045) on: **Monday, September 15, 2014**

Going to close this as it seems to be pretty well solved.
