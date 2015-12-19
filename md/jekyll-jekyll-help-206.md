# [Note block](https://github.com/jekyll/jekyll-help/issues/206)

> state: **open** opened by: **** on: ****

Hi,
I have seen the Jekyll web site and i love the idea to have hint, note warning/ alert box but if i just type markdown note: it does not render like the jekyll site.

how is it possible to have a such box ? also like the ProTip box.

Cheers

### Comments

---
> from: [**alfredxing**](https://github.com/jekyll/jekyll-help/issues/206#issuecomment-65822583) on: **Friday, December 05, 2014**

The notes that appear on the Jekyll documentation site are just snippets of HTML styled with custom CSS:

&#x60;&#x60;&#x60;html
&lt;div class=&quot;note&quot;&gt;
  &lt;h5&gt;ProTips™ help you get more from Jekyll&lt;/h5&gt;
  &lt;p&gt;These are tips and tricks that will help you be a Jekyll wizard!&lt;/p&gt;
&lt;/div&gt;
&#x60;&#x60;&#x60;

&#x60;&#x60;&#x60;scss
.note {
  margin: 30px 0;
  margin-left: -30px;
  padding: 20px 20px 24px;
  padding-left: 50px;
  @include border-radius(0 5px 5px 0);
  position: relative;
  @include box-shadow(0 1px 5px rgba(0, 0, 0, .3), inset 0 1px 0 rgba(255,255,255,.2), inset 0 -1px 0 rgba(0,0,0,.3));
  background-color: #7e6d42;
  background-image: url(data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiA/Pgo8c3ZnIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgd2lkdGg9IjEwMCUiIGhlaWdodD0iMTAwJSIgdmlld0JveD0iMCAwIDEgMSIgcHJlc2VydmVBc3BlY3RSYXRpbz0ibm9uZSI+CiAgPGxpbmVhckdyYWRpZW50IGlkPSJncmFkLXVjZ2ctZ2VuZXJhdGVkIiBncmFkaWVudFVuaXRzPSJ1c2VyU3BhY2VPblVzZSIgeDE9IjAlIiB5MT0iMCUiIHgyPSIwJSIgeTI9IjEwMCUiPgogICAgPHN0b3Agb2Zmc2V0PSIwJSIgc3RvcC1jb2xvcj0iIzdlNmQ0MiIgc3RvcC1vcGFjaXR5PSIxIi8+CiAgICA8c3RvcCBvZmZzZXQ9IjEwMCUiIHN0b3AtY29sb3I9IiM1YzRlMzUiIHN0b3Atb3BhY2l0eT0iMSIvPgogIDwvbGluZWFyR3JhZGllbnQ+CiAgPHJlY3QgeD0iMCIgeT0iMCIgd2lkdGg9IjEiIGhlaWdodD0iMSIgZmlsbD0idXJsKCNncmFkLXVjZ2ctZ2VuZXJhdGVkKSIgLz4KPC9zdmc+);
  background-image: -webkit-gradient(linear, left top, left bottom, from(#7e6d42), to(#5c4e35));
  background-image: -webkit-linear-gradient(top, #7e6d42 0%, #5c4e35 100%);
  background-image: -moz-linear-gradient(top, #7e6d42 0%, #5c4e35 100%);
  background-image: -o-linear-gradient(top, #7e6d42 0%, #5c4e35 100%);
  background-image: linear-gradient(to bottom, #7e6d42 0%,#5c4e35 100%);
  filter: progid:DXImageTransform.Microsoft.gradient( startColorstr=&#x27;#7e6d42&#x27;, endColorstr=&#x27;#5c4e35&#x27;,GradientType=0 );
}
&#x60;&#x60;&#x60;

(full SCSS starting [here](https://github.com/jekyll/jekyll/blob/89bdd47ebcf0e3a60ff0f1795eaabc391fc2307e/site/_sass/_style.scss#L811)).
---
> from: [**zemadmat**](https://github.com/jekyll/jekyll-help/issues/206#issuecomment-65863906) on: **Friday, December 05, 2014**

Thank you for your answer. Instead of using this
&#x60;&#x60;&#x60;
&lt;div class=&quot;note&quot;&gt;
  &lt;h5&gt;ProTips™ help you get more from Jekyll&lt;/h5&gt;
  &lt;p&gt;These are tips and tricks that will help you be a Jekyll wizard!&lt;/p&gt;
&lt;/div&gt;
&#x60;&#x60;&#x60;
is it possible to simply use a markdown syntax ?
---
> from: [**alfredxing**](https://github.com/jekyll/jekyll-help/issues/206#issuecomment-65882935) on: **Friday, December 05, 2014**

I don&#x27;t think so, since Markdown doesn&#x27;t have variables or anything. However, it would be possible to use [includes](http://jekyllrb.com/docs/templates/#includes).
