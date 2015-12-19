# [YAML configuration problem when using &quot;gh-pages&quot;](https://github.com/jekyll/jekyll-help/issues/215)

> state: **closed** opened by: **** on: ****

I&#x27;m getting errors on the settings page when pushing things to github:

&quot;Your page is having problems building: Page build failed. For more information, see https://help.github.com/articles/using-jekyll-with-pages#troubleshooting.&quot;

So I figured my &quot;url&quot; and &quot;baseurl&quot; are wrong in my _config.yml file.

My _config.yml file:

========================================
#Syntax highlighting
highlighter:         pygments

# Permalinks
#Use of &#x60;relative_permalinks&#x60; ensures post links from the index work properly.
permalink:           pretty
relative_permalinks: true

#Setup
title:               Magento design
tagline:          Magento designs and such...


#Local
#http://localhost:4000/magentoproto/homepage-v1/

# This is what I can&#x27;t get to work? Doesn&#x27;t work locally or on github
#url:                 https://github.corp.ebay.com/pages/ajlohman/magentoproto
#baseurl:         /pages/ajlohman/magentoproto


#This works locally but error when pushed to github
#url:                 http://ajlohman.github.io/magentoproto
#baseurl:         /magentoproto


# Assets
#
# We specify the directory for Jekyll so we can use @imports.
sass:
  sass_dir:          _scss
  style:            :compressed

========================================

Any help would be greatly appreciated! thanks,



### Comments

---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/215#issuecomment-67890956) on: **Monday, December 22, 2014**

Hi @ajlohman have seen this? It&#x27;s a good place to start with baseurl issues, as, at a glance, it looks like that might be an issue for you: https://byparker.com/blog/2014/clearing-up-confusion-around-baseurl/

if you want to post a link to your repo, that&#x27;d be helpful too.
---
> from: [**ajlohman**](https://github.com/jekyll/jekyll-help/issues/215#issuecomment-67894939) on: **Monday, December 22, 2014**

Thanks for sending the article!  I still can&#x27;t get it to work, wondering if it&#x27;s something different?

Here&#x27;s the repo: https://github.corp.ebay.com/alohman/magentoproto/tree/gh-pages
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/215#issuecomment-67896034) on: **Monday, December 22, 2014**

So I&#x27;m assuming you changed your config file based on the article? In your second example, you have &#x60;&#x60;&#x60;magentoproto&#x60;&#x60;&#x60; in the url and the base url. 

One thing I might add - I rarely use baseurl any longer. If I need to see a live site for staging I use a subdomain, which is simplistic, but once I have one domain set up, everything else is just a matter of adding a cname file with the subdomain.

is this the repo? https://github.com/ajlohman/magentoproto I don&#x27;t see the config file, but the other looks like it&#x27;s sitting on your corporate domain. 
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/215#issuecomment-68661133) on: **Sunday, January 04, 2015**

I mentioned it [here](https://github.com/jekyll/jekyll-help/issues/226#issuecomment-68659540), but I&#x27;ll mention it again. Use [GitHub Pages Metadata](https://help.github.com/articles/repository-metadata-on-github-pages/). It&#x27;s immensely helpful in situations like this.
