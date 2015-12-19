# [TOC from GFM](https://github.com/jekyll/jekyll-help/issues/298)

> state: **open** opened by: **** on: ****

I would like to have Jekyll generate a table of contents based on headers in a markdown file. I&#x27;ve tried several things based on some research, but no luck. Can anyone point me in the right direction? This is for GitHub Pages.

### Comments

---
> from: [**tomjohnson1492**](https://github.com/jekyll/jekyll-help/issues/298#issuecomment-100275259) on: **Friday, May 08, 2015**

I show how to create a table of contents in this documentation-oriented theme: https://github.com/tomjohnson1492/documentation-theme-jekyll

You use a data file (YML format) with various levels. Then you use a for loop in your theme to iterate over the levels and get the right items in each level. 

I also incorporate the navgoco jquery plugin.
---
> from: [**tonysneed**](https://github.com/jekyll/jekyll-help/issues/298#issuecomment-100484924) on: **Saturday, May 09, 2015**

Hi Tom,

I took a look at your theme, and I very much like it.  My main question would be whether your approach would be compatible with GitHub Pages, as I&#x27;m not sure they include the required plugins.  Would you happen to know?

Cheers,
Tony
---
> from: [**tomjohnson1492**](https://github.com/jekyll/jekyll-help/issues/298#issuecomment-100499734) on: **Saturday, May 09, 2015**

The theme doesn&#x27;t have any plugins. You just need to put one of the configuration files into the root and rename it _config.yml so github pages can build it. 

Sent from my iPhone

&gt; On May 9, 2015, at 6:37 AM, Anthony Sneed &lt;notifications@github.com&gt; wrote:
&gt; 
&gt; Hi Tom,
&gt; 
&gt; I took a look at your theme, and I very much like it. My main question would be whether your approach would be compatible with GitHub Pages, as I&#x27;m not sure they include the required plugins. Would you happen to know?
&gt; 
&gt; Cheers,
&gt; Tony
&gt; 
&gt; —
&gt; Reply to this email directly or view it on GitHub.
&gt; 

---
> from: [**tonysneed**](https://github.com/jekyll/jekyll-help/issues/298#issuecomment-100522749) on: **Saturday, May 09, 2015**

Terrific, I&#x27;ll give it a try!


---
> from: [**tomjohnson1492**](https://github.com/jekyll/jekyll-help/issues/298#issuecomment-100571784) on: **Saturday, May 09, 2015**

Just curious,  are you a technical writer of some kind? What is it about the theme that appeals to you?

Sent from my iPhone

&gt; On May 9, 2015, at 10:47 AM, Anthony Sneed &lt;notifications@github.com&gt; wrote:
&gt; 
&gt; Terrific, I&#x27;ll give it a try!
&gt; 
&gt; —
&gt; Reply to this email directly or view it on GitHub.
&gt; 

---
> from: [**tonysneed**](https://github.com/jekyll/jekyll-help/issues/298#issuecomment-100583661) on: **Saturday, May 09, 2015**

I&#x27;m writing documentation for one of my open source projects on GitHub, and I&#x27;m dissatisfied with the suitability of the default themes for use with a programming reference.

One shortcoming has been the outlining.  For example, I would like to have a numbered list with bulleted items for each number, as well as images and code sections.  But the numbering will restart when followed by bulleted items, and the indenting past the first level didn&#x27;t look right.

So my main question would be how to get your theme to work with GitHub Pages.  I read that you don&#x27;t recommend it for multiple outputs.  But what I like about GitHub Pages, is that I can just push my changes to the repo, and GitHub will automatically build the doc site.  Would it be possible to just have a single level with all my md files in one place and a _config.yml file at the root?  My documentation does not include many pages.

Otherwise, are there other hosting options available where GitHub hooks that can be used to trigger a build and deployment of the doc site after pushing changes to the GitHub repo?
