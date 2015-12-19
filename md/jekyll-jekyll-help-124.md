# [Front Matter Title and filename Title](https://github.com/jekyll/jekyll-help/issues/124)

> state: **closed** opened by: **** on: ****

I am a little confused about Titles... the file/post actual filename serves as a title, but what about the front matter Title? In looking at the Docs page: 
http://jekyllrb.com/docs/frontmatter/ 

the example at the top has Title in the front matter, but down below in the Pre Defined Variables section, Title is not one of them.

I have been using Title in the front matter assuming it was a Pre Defined Variable and would over write whatever the filename was, but I don&#x27;t think it is, and if not should it even be in the front matter?

I noticed this because I have a bunch of posts where the actual filename contains spaces and is capitalized like it should read (rather than lowercase with dashes), but the generated url does not swap out spaces for dashes. Those posts all have Title in the front matter and I was assuming that was the real title as far as Jekyll was concerned,

Should jekyll respect the front matter Title and rename the file with this title, and dashes instead of spaces? if so it isn&#x27;t working.

Is Title in the front matter ignored? if so, the example on http://jekyllrb.com/docs/frontmatter/ should probably be changed as it is a little mis leading.

### Comments

---
> from: [**rdyar**](https://github.com/jekyll/jekyll-help/issues/124#issuecomment-52341100) on: **Friday, August 15, 2014**

I see a little further down that the Title in the front matter was an example of a custom variable, and only does stuff inside the page when called with Liquid as post.title.

This seems a little confusing. If I don&#x27;t put title in the FM does post.title pull in the whole filename  with the date? and dashes?

I&#x27;d rather the FM Title, if it exists becomes the title and over writes the filename, but I see how this is a problem either way.
---
> from: [**rdyar**](https://github.com/jekyll/jekyll-help/issues/124#issuecomment-52342079) on: **Friday, August 15, 2014**

I think I see what is happening, Title in the FM is not necessary as long as you name the file with what you want the title to be... it is just a little confusing, but I think it probably needs to be the way it is.
