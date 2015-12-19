# [Breadcrumb plugin](https://github.com/jekyll/jekyll-help/issues/146)

> state: **open** opened by: **** on: ****

Hello community,

I am looking for a breadcrumb plugin that allow myself to do this : Home / Blog / Category / Title page

I tried several solutions found on the net, but none works :( :

* http://biosphere.cc/software-engineering/jekyll-breadcrumbs-navigation-plugin
* https://github.com/socrata/labs-common-jekyll/blob/master/with-breadcrumbs.html (not the path name)
* http://stackoverflow.com/questions/9612235/what-are-some-good-ways-to-implement-breadcrumbs-on-a-jekyll-site

Do you have a reliable solution ?

Thank you so much :)

### Comments

---
> from: [**liquidvisual**](https://github.com/jekyll/jekyll-help/issues/146#issuecomment-58982945) on: **Monday, October 13, 2014**

Yeah none of these work for me either. Would really like to hear someone&#x27;s solution!
---
> from: [**janvanmansum**](https://github.com/jekyll/jekyll-help/issues/146#issuecomment-59008705) on: **Tuesday, October 14, 2014**

The one from biosphere works for me. 
* Copy-paste the plugin code to &amp;lt;your-project&gt;/_plugins/breadcrumbs.rb
* Create a file &amp;lt;your-project&gt;/_includes/breadcrumbs.html:
&#x60;&#x60;&#x60;&#x60;
       &lt;ol class=&quot;breadcrumbs&quot;&gt;
           {% for p in page.ancestors %}
           &lt;li&gt;&lt;a href=&quot;{{ site.baseurl}}{{ p.url }}&quot;&gt;{{ p.title }}&lt;/a&gt;&lt;/li&gt;
           {% endfor %}
           &lt;li class=&quot;active&quot;&gt;{{ page.title }}&lt;/li&gt;
       &lt;/ol&gt;
&#x60;&#x60;&#x60;&#x60;

* Create a file &amp;lt;your-project&gt;_sass/_breadcrumbs.scss:
&#x60;&#x60;&#x60;&#x60;
    .breadcrumbs {
        margin-left: 0;
        li {
               display: inline;
               list-style: none;
               &amp;:after {
                  content: &quot; &gt; &quot;;
               }
              &amp;.active {
                  &amp;:after {
                      content: &quot;&quot;;
              }
           }
        }
    }
&#x60;&#x60;&#x60;&#x60;
* Add the _breadcrumbs.scss  to &amp;lt;your project&gt;/css/main.scss:
&#x60;&#x60;&#x60;&#x60;
    @import
    &quot;base&quot;,
    &quot;layout&quot;,
    &quot;syntax-highlighting&quot;,
    &quot;breadcrumbs&quot;,
    ....
&#x60;&#x60;&#x60;&#x60;
Then the titles of the index files show up something like this:

Home &gt; Manual &gt; Code Style

if your directory tree is laid out like this:
&#x60;&#x60;&#x60;&#x60;
 [root of project] 
        - index.html  (with &quot;title: Home&quot; in the metadata)
        [manual]
             - index.md (with &quot;title: Manual&quot;  in the metadata)
             [code]
                   - index.md (with &quot;title: Code Style&quot; in the metadata)
&#x60;&#x60;&#x60;&#x60;                  
For posts too many &quot;&gt;&quot; characters appear though.
