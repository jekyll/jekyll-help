# [@Import SASS/Compass Sprites](https://github.com/jekyll/jekyll-help/issues/201)

> state: **closed** opened by: **** on: ****

Hi,

I am struggling to get Jekyll working with Compass Sprites, which uses Sass under the covers.
I am using the jekyll-Assets plugin, and all I want to do is get the site up and running using SASS/Compass.

I am using Compass Sprites.
_assets\images\source\1.png
_assets\images\source\2.png
_assets\images\source\3.png

When I have this in my scss

//Sprites
@import &quot;source/*.png&quot;;


This is easy to reproduce. Just get the latest code from https://github.com/ixti/ixti.github.com. Then add an @import to a subfolder for images and some test images.

I get an error with Jekkyl build.

Liquid exception: No files were found in the load path matching &quot;source/*.png&quot;. You current load paths are ./images ....

How can I get Compass Sprite support work to work correctly?

here is the full stack trace

bsheraton-lc:rangerromblog romikoderbynew$ jekyll build -t
Configuration file: /Users/romikoderbynew/dev/rangerromblog/_config.yml
Source: /Users/romikoderbynew/dev/rangerromblog
Destination: /Users/romikoderbynew/dev/rangerromblog/_site
Generating... 
Liquid Exception: No files were found in the load path matching &quot;source/.png&quot;. Your current load paths are: ./images (in /Users/romikoderbynew/dev/rangerromblog/_assets/stylesheets/application.scss:20) in _layouts/default.html
/Users/romikoderbynew/dev/rangerromblog/_assets/stylesheets/_shared.scss:20: No files were found in the load path matching &quot;source/.png&quot;. Your current load paths are: ./images (Sass::SyntaxError)
(in /Users/romikoderbynew/dev/rangerromblog/_assets/stylesheets/application.scss:20)
from /Users/romikoderbynew/dev/rangerromblog/_assets/stylesheets/application.scss:15
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/tree/import_node.rb:66:in rescue in import&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/tree/import_node.rb:45:inimport&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/tree/import_node.rb:28:in imported_file&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/tree/import_node.rb:37:incss_import?&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/tree/visitors/perform.rb:301:in visit_import&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/tree/visitors/base.rb:36:invisit&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/tree/visitors/perform.rb:158:in block in visit&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/stack.rb:79:inblock in with_base&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/stack.rb:115:in with_frame&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/stack.rb:79:inwith_base&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/tree/visitors/perform.rb:158:in visit&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/tree/visitors/perform.rb:315:inblock (2 levels) in visit_import&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/tree/visitors/perform.rb:315:in map&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/tree/visitors/perform.rb:315:inblock in visit_import&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/stack.rb:88:in block in with_import&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/stack.rb:115:inwith_frame&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/stack.rb:88:in with_import&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/tree/visitors/perform.rb:312:invisit_import&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/tree/visitors/base.rb:36:in visit&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/tree/visitors/perform.rb:158:inblock in visit&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/stack.rb:79:in block in with_base&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/stack.rb:115:inwith_frame&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/stack.rb:79:in with_base&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/tree/visitors/perform.rb:158:invisit&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/tree/visitors/base.rb:52:in block in visit_children&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/tree/visitors/base.rb:52:inmap&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/tree/visitors/base.rb:52:in visit_children&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/tree/visitors/perform.rb:167:inblock in visit_children&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/tree/visitors/perform.rb:179:in with_environment&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/tree/visitors/perform.rb:166:invisit_children&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/tree/visitors/base.rb:36:in block in visit&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/tree/visitors/perform.rb:186:invisit_root&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/tree/visitors/base.rb:36:in visit&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/tree/visitors/perform.rb:157:invisit&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/tree/visitors/perform.rb:8:in visit&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/tree/root_node.rb:36:incss_tree&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/tree/root_node.rb:20:in render&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sass-3.4.8/lib/sass/engine.rb:268:inrender&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sprockets-sass-1.3.0/lib/sprockets/sass/sass_template.rb:53:in evaluate&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/tilt-1.4.1/lib/tilt/template.rb:103:inrender&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sprockets-2.12.3/lib/sprockets/context.rb:197:in block in evaluate&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sprockets-2.12.3/lib/sprockets/context.rb:194:ineach&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sprockets-2.12.3/lib/sprockets/context.rb:194:in evaluate&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sprockets-2.12.3/lib/sprockets/processed_asset.rb:12:ininitialize&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sprockets-2.12.3/lib/sprockets/base.rb:374:in new&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sprockets-2.12.3/lib/sprockets/base.rb:374:inblock in build_asset&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sprockets-2.12.3/lib/sprockets/base.rb:395:in circular_call_protection&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sprockets-2.12.3/lib/sprockets/base.rb:373:inbuild_asset&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sprockets-2.12.3/lib/sprockets/index.rb:94:in block in build_asset&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sprockets-2.12.3/lib/sprockets/caching.rb:51:incache_asset&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sprockets-2.12.3/lib/sprockets/index.rb:93:in build_asset&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sprockets-2.12.3/lib/sprockets/base.rb:287:infind_asset&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sprockets-2.12.3/lib/sprockets/index.rb:61:in find_asset&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/jekyll-assets-0.11.0/lib/jekyll/assets_plugin/patches/index_patch.rb:16:in__wrap_find_asset&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sprockets-2.12.3/lib/sprockets/bundled_asset.rb:16:in initialize&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sprockets-2.12.3/lib/sprockets/base.rb:377:innew&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sprockets-2.12.3/lib/sprockets/base.rb:377:in build_asset&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sprockets-2.12.3/lib/sprockets/index.rb:94:inblock in build_asset&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sprockets-2.12.3/lib/sprockets/caching.rb:51:in cache_asset&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sprockets-2.12.3/lib/sprockets/index.rb:93:inbuild_asset&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sprockets-2.12.3/lib/sprockets/base.rb:287:in find_asset&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sprockets-2.12.3/lib/sprockets/index.rb:61:infind_asset&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/jekyll-assets-0.11.0/lib/jekyll/assets_plugin/patches/index_patch.rb:16:in __wrap_find_asset&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sprockets-2.12.3/lib/sprockets/environment.rb:75:infind_asset&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/jekyll-assets-0.11.0/lib/jekyll/assets_plugin/environment.rb:52:in find_asset&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/sprockets-2.12.3/lib/sprockets/base.rb:295:in[]&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/jekyll-assets-0.11.0/lib/jekyll/assets_plugin/renderer.rb:53:in render_tag&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/jekyll-assets-0.11.0/lib/jekyll/assets_plugin/renderer.rb:39:inrender_stylesheet&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/jekyll-assets-0.11.0/lib/jekyll/assets_plugin/tag.rb:11:in render&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/liquid-2.6.1/lib/liquid/block.rb:109:inblock in render_all&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/liquid-2.6.1/lib/liquid/block.rb:96:in each&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/liquid-2.6.1/lib/liquid/block.rb:96:inrender_all&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/liquid-2.6.1/lib/liquid/block.rb:82:in render&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/liquid-2.6.1/lib/liquid/template.rb:128:inrender&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/liquid-2.6.1/lib/liquid/template.rb:138:in render!&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/jekyll-2.4.0/lib/jekyll/convertible.rb:106:inrender_liquid&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/jekyll-2.4.0/lib/jekyll/convertible.rb:205:in render_all_layouts&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/jekyll-2.4.0/lib/jekyll/convertible.rb:239:indo_layout&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/jekyll-2.4.0/lib/jekyll/post.rb:261:in render&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/jekyll-2.4.0/lib/jekyll/site.rb:269:inblock in render&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/jekyll-2.4.0/lib/jekyll/site.rb:268:in each&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/jekyll-2.4.0/lib/jekyll/site.rb:268:inrender&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/jekyll-2.4.0/lib/jekyll/site.rb:46:in process&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/jekyll-2.4.0/lib/jekyll/command.rb:28:inprocess_site&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/jekyll-2.4.0/lib/jekyll/commands/build.rb:55:in build&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/jekyll-2.4.0/lib/jekyll/commands/build.rb:33:inprocess&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/jekyll-2.4.0/lib/jekyll/commands/build.rb:17:in block (2 levels) in init_with_program&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/mercenary-0.3.5/lib/mercenary/command.rb:220:incall&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/mercenary-0.3.5/lib/mercenary/command.rb:220:in block in execute&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/mercenary-0.3.5/lib/mercenary/command.rb:220:ineach&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/mercenary-0.3.5/lib/mercenary/command.rb:220:in execute&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/mercenary-0.3.5/lib/mercenary/program.rb:42:ingo&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/mercenary-0.3.5/lib/mercenary.rb:19:in program&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/jekyll-2.4.0/bin/jekyll:18:in&#x27;
from /Users/romikoderbynew/.rbenv/versions/2.1.3/bin/jekyll:23:in load&#x27; from /Users/romikoderbynew/.rbenv/versions/2.1.3/bin/jekyll:23:in&#x27;

### Comments

---
> from: [**sondr3**](https://github.com/jekyll/jekyll-help/issues/201#issuecomment-65016052) on: **Sunday, November 30, 2014**

Curious, have you tried trying to import them with the full path? Seems like SASS is confused about the directories, if you tried something like &#x60;&#x60;&#x60;@import &#x27;~/dev/rangerromblog/_assets/images/*.png&#x27;&#x60;&#x60;&#x60;. I haven&#x27;t really used Jekyll-Assets or anything like that since I use Gulp myself for this.
---
> from: [**Romiko**](https://github.com/jekyll/jekyll-help/issues/201#issuecomment-65016121) on: **Sunday, November 30, 2014**

Hi, I just use compass compile and removed jekyll assets.  Then just wrote
my own plugin to combine asaets.
On 01/12/2014 2:12 PM, &quot;Sondre Nilsen&quot; &lt;notifications@github.com&gt; wrote:

&gt; Curious, have you tried trying to import them with the full path? Seems
&gt; like SASS is confused about the directories, if you tried something like @import
&gt; &#x27;~/dev/rangerromblog/_assets/images/*.png&#x27;. I haven&#x27;t really used
&gt; Jekyll-Assets or anything like that since I use Gulp myself for this.
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub
&gt; &lt;https://github.com/jekyll/jekyll-help/issues/201#issuecomment-65016052&gt;.
&gt;
---
> from: [**sondr3**](https://github.com/jekyll/jekyll-help/issues/201#issuecomment-65016376) on: **Sunday, November 30, 2014**

Alright, good to know you solved it! You can probably close this issue as it&#x27;s no longer a problem for you :)
