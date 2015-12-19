# [category[0] is all lower case while tag[0] is case sensitive?](https://github.com/jekyll/jekyll-help/issues/125)

> state: **open** opened by: **** on: ****

I&#x27;m working up a page that lists all posts by tag, and another page is all posts by category.

I am getting category[0] returning the category name as all lowercase (even if it is Capitalized in the front matter) while tag[0] returns what I put into the front matter.

My front matter would look like:

&#x60;&#x60;&#x60;
---
category: Jekyll Update
tags: [Jekyll]
---
&#x60;&#x60;&#x60;

This seems wrong. I can capitalize with liquid, but it seems like it should be consistent.

It is possible that I am not using the category right I suppose. I am pretty confused by how it works sometimes. I am not using folders, other than to organize - right under the posts folder like:
_posts
          videos
          articles

i consider videos and articles to be types - not categories or tags and handle them myself with custom front matter.

### Comments

