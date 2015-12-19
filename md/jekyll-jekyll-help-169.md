# [Cannot load rdiscount.so (markdown processing)](https://github.com/jekyll/jekyll-help/issues/169)

> state: **open** opened by: **** on: ****

I&#x27;m trying to use Octopress to generate my first post, but maybe this is something Jekyll can answer as it&#x27;s used as the basis for Octopress.. Still can&#x27;t get it to process the markdown post..

I&#x27;m running Windows 7 64 bits, installed Ruby 1.9.3 and the latest Octopress from the git repository (from today).

I&#x27;ve verified that rdiscount is correctly installed (through the bundle install).
rdiscount is 2.1.7.1. rdiscount.so is in ext/rdiscount.so..

Anybody has an idea of what is going wrong ? Thanks.

&gt; rake generate
## Set the codepage to 65001 for Windows machines
## Generating Site with Jekyll
identical source/stylesheets/screen.css
Configuration file: C:/website/octopress/_config.yml
            Source: source
       Destination: ../production/public/
      Generating...
[31m  Dependency Error:  Yikes! It looks like you don&#x27;t have rdiscount or one o
f its dependencies installed. In order to use Jekyll as currently configured, yo
u&#x27;ll need to install this gem. The full error message from Ruby is: &#x27;cannot load
 such file -- rdiscount.so&#x27; If you run into trouble, you can find helpful
 resources at http://jekyllrb.com/help/! [0m
[31m  Conversion error: Jekyll::Converters::Markdown encountered an error conve
rting &#x27;_posts/2008-10-23-tt.md/#excerpt&#x27;.[0m
[31m  Conversion error: rdiscount[0m
[31m             ERROR: YOUR SITE COULD NOT BE BUILT:[0m
[31m                    ------------------------------------[0m
[31m                    rdiscount[0m

### Comments

---
> from: [**thinkinglk**](https://github.com/jekyll/jekyll-help/issues/169#issuecomment-60059307) on: **Wednesday, October 22, 2014**

maybe you should install rdiscount,  use the command &quot;gem install rdiscount&quot;
