# [category_generator sort! error](https://github.com/jekyll/jekyll-help/issues/204)

> state: **open** opened by: **** on: ****

The jekyll build processes died tonight on categories.sort!.map { |c| category_link c } when processing the YAML head of templates for posts. I commented out the lines in plugins/category_generator.rb and the site generate process worked, but the YAML data was turned into p of text.

I later updated bundler and was able to restore the template processing modules. The failure to process seems to have been caused by titles were not quoted. After fixing this, the posts at least were processed.

Processed YAML head was turned into a single paragraph for one file, then two files, then all files. Very strange that not all went at once.

It turns out the rake routine that builds posts adds a time zone number (-0800) to the end of the date. In my rake file for drafts-&gt;posts I included this, but instead of using z% I added the number it would generate, the -0800. When I deleted the time zone the processing of all YAML head material worked fine. Seems a little picky, eh?

### Comments

---
> from: [**sondr3**](https://github.com/jekyll/jekyll-help/issues/204#issuecomment-65015846) on: **Sunday, November 30, 2014**

Anywhere to take a look at the source code? This kinda looks like there&#x27;s some kind of formatting error or something like that in one of your layouts, I have had similar errors occasionally when I mess around with Jekyll too.
