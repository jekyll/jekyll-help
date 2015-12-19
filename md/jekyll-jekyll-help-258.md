# [Pagination Duplicating Root Index.html](https://github.com/jekyll/jekyll-help/issues/258)

> state: **open** opened by: **** on: ****

I&#x27;m trying to follow the instructions of pagination on Jekyll. However, after editing the _config.yml file and rebuilding my site, my root index.html file is duplicated within the pagination directory.

The Jekyll site has the following layout:
_data
_includes
_layouts
_posts
_site
css
about.html
authors.html
hof.html
index.html
issues.html
links.html
news.html

Inside _site the site looks like this:
about
css
haymaker
issues
about.html
authors.html
hof.html
index.html
issues.html
links.html
news.html

The problem is that in _posts, I have 66 files. I want have 10 items per page in the issues directory. We took the issues.html file, renamed it index.html and dropped it into the issues directory, and then had Jekyll rebuild the site. However, once I rebuilt the site and then looked at the directory structure, the index.html file had vanished from the issues directory, and while I had folders labeled page 2-7, each of them contained just one file, which was a duplicate of the index.html from my root directory.

I have a screenshot here:
![directory_strucutre](https://cloud.githubusercontent.com/assets/3038283/5970271/df1d54c2-a7fe-11e4-9848-d056a9586a3a.png)

I suspect the issue is with the code I am using, but I&#x27;m not versed enough with the syntax to figure it out on my own.

Any suggestions?



### Comments

