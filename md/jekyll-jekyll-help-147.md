# [Pagination not working as expected](https://github.com/jekyll/jekyll-help/issues/147)

> state: **open** opened by: **** on: ****

I am trying to use pagination, but being new to all of this doesn&#x27;t help and looking for answers here and on Google are proving itself a challenge.

All I would like to do is split up my blog to separate pages.

I have this so far on my index:

{% for post in paginator.posts %}
--Post Title--
{% endfor %}

And the paginate variable is set to 5 in my config.


### Comments

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/147#issuecomment-55484533) on: **Saturday, September 13, 2014**

Could you describe what exactly isn not working as expected? What do you get and what did you expect?

Did you follow the [docs about pagination](http://jekyllrb.com/docs/pagination/#render-the-paginated-posts)? The examples provided there should work. If they do not, come back and tell us. :)
