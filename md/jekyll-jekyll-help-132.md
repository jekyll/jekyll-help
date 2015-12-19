# [SASS won’t compile with jekyll build](https://github.com/jekyll/jekyll-help/issues/132)

> state: **open** opened by: **** on: ****

On [this](https://github.com/ndarville/jekyll-boilerplate/blob/699840d35521e01c79a4c2d4a9e13887ede120f2/static/css/style.scss) project (YAML header included in local version), I get the error below, when running &#x60;jekyll build --trace&#x60;.

I have tried appending &#x60;.scss&#x60; to the imports, but I get the same error.

&#x60;&#x60;&#x60;sh
$ jekyll build --trace
Notice: for 10x faster LSI support, please install http://rb-gsl.rubyforge.org/
Configuration file: none
            Source: /jekyll-boilerplate/static/css
       Destination: /jekyll-boilerplate/static/css/_site
      Generating... 
  Conversion error: There was an error converting &#x27;style.scss&#x27;.
(sass):1: File to import not found or unreadable: _reset. (Sass::SyntaxError)
&#x60;&#x60;&#x60;

&#x60;&#x60;&#x60;sh
Load path: /jekyll-boilerplate/static/css/_sass
    from /.rvm/gems/ruby-2.1.0/gems/sass-3.3.14/lib/sass/tree/import_node.rb:66:in &#x60;rescue in import&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/sass-3.3.14/lib/sass/tree/import_node.rb:45:in &#x60;import&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/sass-3.3.14/lib/sass/tree/import_node.rb:28:in &#x60;imported_file&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/sass-3.3.14/lib/sass/tree/import_node.rb:37:in &#x60;css_import?&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/sass-3.3.14/lib/sass/tree/visitors/perform.rb:283:in &#x60;visit_import&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/sass-3.3.14/lib/sass/tree/visitors/base.rb:36:in &#x60;visit&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/sass-3.3.14/lib/sass/tree/visitors/perform.rb:152:in &#x60;block in visit&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/sass-3.3.14/lib/sass/stack.rb:79:in &#x60;block in with_base&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/sass-3.3.14/lib/sass/stack.rb:121:in &#x60;with_frame&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/sass-3.3.14/lib/sass/stack.rb:79:in &#x60;with_base&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/sass-3.3.14/lib/sass/tree/visitors/perform.rb:152:in &#x60;visit&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/sass-3.3.14/lib/sass/tree/visitors/base.rb:52:in &#x60;block in visit_children&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/sass-3.3.14/lib/sass/tree/visitors/base.rb:52:in &#x60;map&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/sass-3.3.14/lib/sass/tree/visitors/base.rb:52:in &#x60;visit_children&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/sass-3.3.14/lib/sass/tree/visitors/perform.rb:161:in &#x60;block in visit_children&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/sass-3.3.14/lib/sass/tree/visitors/perform.rb:173:in &#x60;with_environment&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/sass-3.3.14/lib/sass/tree/visitors/perform.rb:160:in &#x60;visit_children&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/sass-3.3.14/lib/sass/tree/visitors/base.rb:36:in &#x60;block in visit&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/sass-3.3.14/lib/sass/tree/visitors/perform.rb:180:in &#x60;visit_root&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/sass-3.3.14/lib/sass/tree/visitors/base.rb:36:in &#x60;visit&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/sass-3.3.14/lib/sass/tree/visitors/perform.rb:151:in &#x60;visit&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/sass-3.3.14/lib/sass/tree/visitors/perform.rb:8:in &#x60;visit&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/sass-3.3.14/lib/sass/tree/root_node.rb:36:in &#x60;css_tree&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/sass-3.3.14/lib/sass/tree/root_node.rb:20:in &#x60;render&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/sass-3.3.14/lib/sass/engine.rb:274:in &#x60;render&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/sass-3.3.14/lib/sass.rb:62:in &#x60;compile&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/jekyll-sass-converter-1.2.0/lib/jekyll/converters/scss.rb:93:in &#x60;convert&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/jekyll-2.2.0/lib/jekyll/convertible.rb:65:in &#x60;transform&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/jekyll-2.2.0/lib/jekyll/convertible.rb:212:in &#x60;do_layout&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/jekyll-2.2.0/lib/jekyll/page.rb:122:in &#x60;render&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/jekyll-2.2.0/lib/jekyll/site.rb:261:in &#x60;block in render&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/jekyll-2.2.0/lib/jekyll/site.rb:260:in &#x60;each&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/jekyll-2.2.0/lib/jekyll/site.rb:260:in &#x60;render&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/jekyll-2.2.0/lib/jekyll/site.rb:43:in &#x60;process&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/jekyll-2.2.0/lib/jekyll/command.rb:53:in &#x60;process_site&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/jekyll-2.2.0/lib/jekyll/commands/build.rb:50:in &#x60;build&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/jekyll-2.2.0/lib/jekyll/commands/build.rb:33:in &#x60;process&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/jekyll-2.2.0/lib/jekyll/commands/build.rb:17:in &#x60;block (2 levels) in init_with_program&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/mercenary-0.3.4/lib/mercenary/command.rb:220:in &#x60;call&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/mercenary-0.3.4/lib/mercenary/command.rb:220:in &#x60;block in execute&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/mercenary-0.3.4/lib/mercenary/command.rb:220:in &#x60;each&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/mercenary-0.3.4/lib/mercenary/command.rb:220:in &#x60;execute&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/mercenary-0.3.4/lib/mercenary/program.rb:35:in &#x60;go&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/mercenary-0.3.4/lib/mercenary.rb:22:in &#x60;program&#x27;
    from /.rvm/gems/ruby-2.1.0/gems/jekyll-2.2.0/bin/jekyll:18:in &#x60;&lt;top (required)&gt;&#x27;
    from /.rvm/gems/ruby-2.1.0/bin/jekyll:23:in &#x60;load&#x27;
    from /.rvm/gems/ruby-2.1.0/bin/jekyll:23:in &#x60;&lt;main&gt;&#x27;
    from /.rvm/gems/ruby-2.1.0/bin/ruby_executable_hooks:15:in &#x60;eval&#x27;
    from /.rvm/gems/ruby-2.1.0/bin/ruby_executable_hooks:15:in &#x60;&lt;main&gt;&#x27;
&#x60;&#x60;&#x60;

The relevant part of &#x60;_config.yml&#x60;:

&#x60;&#x60;&#x60;yaml
sass:
    sass_dir: static/css/
    style: :nested
&#x60;&#x60;&#x60;

### Comments

---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/132#issuecomment-52763611) on: **Wednesday, August 20, 2014**

Maybe it’s because the partials have an underscore. Gonna check it out.
---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/132#issuecomment-52763784) on: **Wednesday, August 20, 2014**

Ah, I see; the config is looking in &#x60;_sass&#x60;. Maybe we can add a config for a &#x60;sass_partial_dir&#x60;?

Using &#x60;_sass&#x60; as a dir for partial SASS, when you have your main SASS in the parent folder, doesn’t makes a lot of semantic sense to me.
---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/132#issuecomment-52764545) on: **Wednesday, August 20, 2014**

So, I got it working with these steps:

- [x] Inserted YAML header in main SCSS files
- [x] Moved partials to &#x60;_sass&#x60;
- [x] Renamed &#x60;@import&#x60; statements to from &#x60;foo&#x60; to &#x60;_sass/foo&#x60;
---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/132#issuecomment-52774207) on: **Wednesday, August 20, 2014**

- [x] Deleted legacy &#x60;.css&#x60; and &#x60;_site/*/.scss&#x60; files

(And added a YAML header that was missing for &#x60;style.scss&#x60; for some reason.)

ndarville/jekyll-boilerplate@0f2332c5de6b0d82fab578a93a4b8f3acd2f31ad
---
> from: [**vivek2014**](https://github.com/jekyll/jekyll-help/issues/132#issuecomment-54060914) on: **Monday, September 01, 2014**


C:\Users\coder\Desktop\Vivek&gt;Jekyll serve
Configuration file: C:/Users/coder/Desktop/Vivek/_config.yml
            Source: C:/Users/coder/Desktop/Vivek
       Destination: C:/Users/coder/Desktop/Vivek/_site
      Generating...
  Conversion error: Jekyll::Converters::Sass encountered an error converting &#x27;as
sets/css/main.sass&#x27;.
jekyll 2.3.0 | Error:  The line was indented 2 levels deeper than the previous l
ine.

Some 1 Please Help What To Do?
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/132#issuecomment-54083840) on: **Monday, September 01, 2014**

@vivek2014 can you post the code? It seems to have a wrong indentation.
