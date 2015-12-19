# [Redirect users to talk.jekyllrb.com](https://github.com/jekyll/jekyll-help/pull/304)

> state: **open** opened by: **** on: ****

Suggested rewrite to begin redirecting users to the new forum. Discussion here:

https://talk.jekyllrb.com/t/retire-jekyll-help-in-favor-of-the-new-forum/170/8?u=erlend_sh

### Comments

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/pull/304#issuecomment-100701337) on: **Sunday, May 10, 2015**

I&#x27;m :+1: on this. How do we migrate the issues, you think?
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/pull/304#issuecomment-100702491) on: **Sunday, May 10, 2015**

We&#x27;d want to either make the comment via the user if one exists with the same email/username or we&#x27;d want to just have system do something like gopherbot does for golang/go:

&#x60;&#x60;&#x60;text
[@erlend-sh](https://github.com/erlend-sh) wrote on 2015-05-10 at 22:03:23 UTC:

&gt; Suggested rewrite to begin redirecting users to the new forum. Discussion here:
&gt;
&gt; https://talk.jekyllrb.com/t/retire-jekyll-help-in-favor-of-the-new-forum/170/8?u=erlend_sh
&#x60;&#x60;&#x60;

which would render as

[@erlend-sh](https://github.com/erlend-sh) wrote on 2015-05-10 at 22:03:23 UTC:

&gt; Suggested rewrite to begin redirecting users to the new forum. Discussion here:
&gt;
&gt; https://talk.jekyllrb.com/t/retire-jekyll-help-in-favor-of-the-new-forum/170/8?u=erlend_sh

and we&#x27;re good.

And do that for each thread. I use [offline-issues](https://github.com/jlord/offline-issues) already to work on Jekyll stuff on the go :airplane: so maybe we can use that to grab the issues and go from there.
---
> from: [**erlend-sh**](https://github.com/jekyll/jekyll-help/pull/304#issuecomment-100735215) on: **Sunday, May 10, 2015**

&gt; How do we migrate the issues, you think?
&gt; (...)
&gt; or we&#x27;d want to just have system do something like gopherbot does for golang/go

That sounds like the way to go to me. Trying to maintain post ownership quickly gets complicated: There&#x27;s 300 issues totalling roughly 600 posts, and I&#x27;m guessing far less than 50% of the posters have a forum account they can be tied to.

offline-issues looks neat. I had a similar idea and found these:

http://jaredly.github.io/github-issues-viewer/#jekyll/jekyll-help
https://github.com/jaxbot/github-issues.vim

I can&#x27;t do anything with any of those myself though.

If you can get me plain-text (i.e. Markdown friendly) transcripts of the 300 issues, I wouldn&#x27;t mind pretending to be your bot for the week, manually posting those issues over the course of a couple of days (e.g. as a new &quot;github-help&quot; user).
