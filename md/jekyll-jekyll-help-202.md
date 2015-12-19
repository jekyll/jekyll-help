# [Using a partial in a markdown file - should I need to wrap it in HTML? ](https://github.com/jekyll/jekyll-help/issues/202)

> state: **closed** opened by: **** on: ****

When including an HTML partial in the middle of a markdown file it seems that I need to wrap it in a div. Am I doing something wrong?

This doesn&#x27;t work, but I&#x27;d expect it to:

&#x60;{% include call-to-action.html %}&#x60;

This does work, but I need to wrap it up in some HTML:

&#x60;&lt;div&gt;{% include call-to-action.html %}&lt;/div&gt;&#x60;

### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/202#issuecomment-64262102) on: **Monday, November 24, 2014**

You don&#x27;t have to. It really depends on what&#x27;s in &#x60;call-to-action.html&#x60;. All the &#x60;{% include ... %}&#x60; tag does is literally include the contents of the file into the layout.

Some more context would be helpful if you need some help.
---
> from: [**barryclark**](https://github.com/jekyll/jekyll-help/issues/202#issuecomment-64396239) on: **Tuesday, November 25, 2014**

Thanks for the help @troyswanson -- I had an indent in my html file that was causing the markup to output incorrectly when pulled into the markdown file. Fixed! :smiley: 
