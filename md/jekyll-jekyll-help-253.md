# [&#x60;:title&#x60; in collections seems to parse punctuation differently](https://github.com/jekyll/jekyll-help/issues/253)

> state: **closed** opened by: **** on: ****

I use a collection to manage “draft” entries, so I can print out a schedule of upcoming entries on my site. But when generating output files from that collection, it seems &#x60;:title&#x60; might be parsed differently in collections than it is elsewhere?

Here’s what I’m working from:

- Post’s filename: &#x60;_collection/2015-01-26-laphams-quarterly.markdown&#x60;
- Post’s title, in the front matter: &#x60;Lapham’s Quarterly”
- Permalink definition for the collection: &#x60;dir/:title:output_ext&#x60;
- Default permalink setting: &#x60;dir/:title:output_ext&#x60;

When served, this results in:

- &#x60;dir/lapham-s-quarterly.html&#x60; if the post is in the collection; and
- &#x60;dir/laphams-quarterly.html&#x60; if the post is in &#x60;_posts&#x60;.

(I am very likely doing something horribly wrong, but thought I’d log it here in case the issue’s relevant to anyone else.)

### Comments

---
> from: [**paulrobertlloyd**](https://github.com/jekyll/jekyll-help/issues/253#issuecomment-71488303) on: **Monday, January 26, 2015**

I suspect this may be fixed when posts become a type of collection… looks like there are some under-the-hood differences in how URL slugs are handled for each type… but that’s just a guess. Am I right @parkr?

By the way @beep, have you seen the built in handling for draft posts: http://jekyllrb.com/docs/drafts/ ?
---
> from: [**beep**](https://github.com/jekyll/jekyll-help/issues/253#issuecomment-71488679) on: **Monday, January 26, 2015**

@paulrobertlloyd I have, thanks! (I opted for a collection because we _do_ want to output files from the “drafts” collection, and to print out a list of the upcoming entries; near as I could tell, those things aren’t possible with &#x60;_drafts&#x60;.)
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/253#issuecomment-71516475) on: **Monday, January 26, 2015**

Hey, @beep! Sorry you ran into this problemo. Turns out we [extract the slug (used by &#x60;:title&#x60;) from the post filename](https://github.com/jekyll/jekyll/blob/master/lib/jekyll/post.rb#L170), whereas collection documents&#x27; slugs [are generated either from the &#x60;title&#x60; element in the YAML front matter or, if it doesn&#x27;t exist, by the file basename](https://github.com/jekyll/jekyll/blob/master/lib/jekyll/document.rb#L135).

It would make sense to me to call the slug the input file&#x27;s basename in all cases (with date data stripped, of course). What say you?
---
> from: [**beep**](https://github.com/jekyll/jekyll-help/issues/253#issuecomment-71516815) on: **Monday, January 26, 2015**

&gt; It would make sense to me to call the slug the input file&#x27;s basename in all cases (with date data stripped, of course). What say you?

@parkr That’d be what I expect, and gets a big +1 from me! Thanks!

(And thanks for weighing in!)
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/253#issuecomment-71773638) on: **Tuesday, January 27, 2015**

Terrific! Would you mind filing a bug for this in jekyll/jekyll? We&#x27;ll get it fixed as soon as we can. :)
---
> from: [**beep**](https://github.com/jekyll/jekyll-help/issues/253#issuecomment-71861795) on: **Wednesday, January 28, 2015**

@parkr Done and done—please let me know if [the bug](https://github.com/jekyll/jekyll/issues/3372) should be worded differently. And thank you so much for weighing in!
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/253#issuecomment-71959620) on: **Wednesday, January 28, 2015**

Tracking over in jekyll/jekyll. Thanks, all!
