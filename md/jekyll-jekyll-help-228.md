# [Build problem - _config.yml format](https://github.com/jekyll/jekyll-help/issues/228)

> state: **closed** opened by: **** on: ****

Hi

I&#x27;m starting my blog site from scratch (I cloned &#x60;poole/hyde&#x60;) and I&#x27;m am trying to test the site locally using &#x60;bundle exec jekyll build&#x60; (I have added a Gemfile to the root of the source folder).  Now I&#x27;m getting the following error:

I am trying to build and test the site locally, using &#x60;bundle exec jekyll build&#x60;
(I’ve added a Gemfile to the root of the source folder) but I’m now getting
the following error:

    $bundle exec jekyll build --watch
    jekyll 2.4.0 | Error:  (/Users/srm/Documents/sandeep/cst/content/sites/sr-murthy.github.io/_config.yml): did not find expected key while parsing a block mapping at line 3 column 1

The first three lines of &#x60;_config.yml&#x60; are

    --- 
    # Site setup, content, dependencies and global properties
    title: &#x27;free to code&#x27;

Please help.




### Comments

---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/228#issuecomment-68930495) on: **Tuesday, January 06, 2015**

Remove the first line &#x60;---&#x60; and try again.
---
> from: [**sr-murthy**](https://github.com/jekyll/jekyll-help/issues/228#issuecomment-68960507) on: **Tuesday, January 06, 2015**

Thanks, I had fixed it yesterday.
