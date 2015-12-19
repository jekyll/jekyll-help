# [Using _config file to rename filenames but retain paths](https://github.com/jekyll/jekyll-help/issues/140)

> state: **open** opened by: **** on: ****

Is there a way I can set the default permalink in my _config file to rename the HTML files for my site while retaining the existing directory structure of my Markdown project? Basically, my source project has a bunch of folders with Markdown files named readme.md (so they will be automatically displayed in GitHub when viewing the folder) like this:

- docs/readme.md
- docs/help/readme.md
- download/readme.md

But for my site, I&#x27;d like to rename them all index.html and retain the exisiting directory structure like this:

- docs/index.html
- docs/help/index.html
- download/index.html

It doesn&#x27;t seem like I can rename just the file. I&#x27;m currently using a permalink variable in each MD file to set the path and the filename for the HTML on the site, but it would be much easier to maintain if I could just set the filename from the config file and tell it to keep the current path. Something like:

*/:dir/index.html* or */:path/index.html*

### Comments

---
> from: [**joewils**](https://github.com/jekyll/jekyll-help/issues/140#issuecomment-53626832) on: **Wednesday, August 27, 2014**

i don&#x27;t think you can have Jekyll rename files for you.  as a work around you could create an index.html file with an include statement to your readme file: {% include readme.md %} 
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/140#issuecomment-53645813) on: **Wednesday, August 27, 2014**

This would make for a great plugin, especially for Jekyll 3.0!
