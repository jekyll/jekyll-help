# [Problem running generator on Windows Server 2008 R2](https://github.com/jekyll/jekyll-help/issues/121)

> state: **closed** opened by: **** on: ****

Hello,

I am hoping that someone can help.

I originally posted this question here:

https://github.com/jekyll/jekyll/issues/2738

but it was suggested that I should ask it here instead.

As requested there by @parkr I have confirmed the jekyll version that is running, and it is 

&#x60;&#x60;&#x60;jekyll 2.2.0&#x60;&#x60;&#x60;

As I think I mentioned below, I have set this up on a fresh Azure VM, which is essentially throw away.  If it will help, and if someone has some time, I can gladly provide access to the VM via RDP in order to help figure out what could be going on.

Original question follows....

Thanks!

-------------------------------------------------

Can anyone help with this issue that I am facing?

I originally posted on the generator-jekyllrb issue tracker, but @robwierzbowski suggested that I post here instead.  The details of the issue are as follows (which was originally posted here https://github.com/robwierzbowski/generator-jekyllrb/issues/99#issuecomment-51968568):

--------------------------------------

I have been trying to use the jekyllrb generator to template out a new Jekyll site I am making.  I have followed the guide here:

http://jekyll-windows.juthilo.com/

And the instructions on this site to try to get it up and running, but I keep getting an error.

These are the steps that I have followed:

 * Install Chocolatey
 * Install notepad++ &#x60;&#x60;&#x60;choco install notepadplusplus&#x60;&#x60;&#x60;
 * Install Ruby
&#x60;&#x60;&#x60; choco install ruby &#x60;&#x60;&#x60;
 * Confirm ruby is installed correctly &#x60;&#x60;&#x60;ruby -v&#x60;&#x60;&#x60; which gives &#x60;&#x60;&#x60;ruby 2.0.0p481&#x60;&#x60;&#x60;
 * Install Ruby DevKit &#x60;&#x60;&#x60;choco install ruby2.devkit&#x60;&#x60;&#x60;
 * Fixed up config.yaml file and then ran &#x60;&#x60;&#x60;ruby dk.rb install&#x60;&#x60;&#x60;
 * Install jekyll &#x60;&#x60;&#x60;gem install jekyll&#x60;&#x60;&#x60;
 * Install rouge &#x60;&#x60;&#x60;gem install rouge&#x60;&#x60;&#x60;
 * Install python &#x60;&#x60;&#x60;choco install python2&#x60;&#x60;&#x60;
 * Confirm python is installed correctly &#x60;&#x60;&#x60;python -V&#x60;&#x60;&#x60; which gives &#x60;&#x60;&#x60;Python 2.7.8&#x60;&#x60;&#x60;
 * Install pip &#x60;&#x60;&#x60;choco install pip&#x60;&#x60;&#x60;
 * Install pygments &#x60;&#x60;&#x60;python -m pip install Pygments&#x60;&#x60;&#x60;
 * Install wdm &#x60;&#x60;&#x60;gem install wdm&#x60;&#x60;&#x60;
 * Install node &#x60;&#x60;&#x60;choco install nodejs.install&#x60;&#x60;&#x60;

Then I tested to see that things are working as expecte

 * &#x60;&#x60;&#x60;jekyll new myblog&#x60;&#x60;&#x60;
 * &#x60;&#x60;&#x60;jekyll serve&#x60;&#x60;&#x60;

All good!

Now onto installing the generator

 * Install bundler &#x60;&#x60;&#x60;gem install bundler&#x60;&#x60;&#x60;
 * Install generator &#x60;&#x60;&#x60;npm install -g generator-jekyllrb&#x60;&#x60;&#x60;
 * And then templated out a site, accepting the default options
 * &#x60;&#x60;&#x60;yo jekyllrb&#x60;&#x60;&#x60;

And I get this at the end:

![image](https://cloud.githubusercontent.com/assets/1271146/3896503/eb0484bc-2258-11e4-91bb-234aa96dd845.png)

So then I did this:

 * &#x60;&#x60;&#x60;bower install&#x60;&#x60;&#x60;
 * &#x60;&#x60;&#x60;npm install&#x60;&#x60;&#x60;
 * Then I noticed errors about missing git so I did &#x60;&#x60;&#x60;choco install git&#x60;&#x60;&#x60;
 * &#x60;&#x60;&#x60;bower install&#x60;&#x60;&#x60;
 * &#x60;&#x60;&#x60;npm install&#x60;&#x60;&#x60;

Then I do this:

&#x60;&#x60;&#x60; grunt serve&#x60;&#x60;&#x60;

And this is what I get:

![image](https://cloud.githubusercontent.com/assets/1271146/3896606/e83e03ec-2259-11e4-9b22-052fd5fa3381.png)

Any ideas what I did wrong?

Thanks!

### Comments

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/121#issuecomment-52012889) on: **Tuesday, August 12, 2014**

helllo again! Did you run &#x60;jekyll build --trace&#x60; to get the full trace output?
---
> from: [**gep13**](https://github.com/jekyll/jekyll-help/issues/121#issuecomment-52013064) on: **Tuesday, August 12, 2014**

Hi!

Here is the full output:

&#x60;&#x60;&#x60;
Notice: for 10x faster LSI support, please install http://rb-gsl.rubyforge.org/
       Deprecation: The &#x27;pygments&#x27; configuration option has been renamed to &#x27;hig
hlighter&#x27;. Please update your config file accordingly. The allowed values are &#x27;r
ouge&#x27;, &#x27;pygments&#x27; or null.
C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/site.rb:133:in
 &#x60;open&#x27;: No such file or directory - . (Errno::ENOENT)
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:133:in &#x60;entries&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:133:in &#x60;block in read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:133:in &#x60;chdir&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:133:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:144:in &#x60;block in read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:144:in &#x60;block in read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:144:in &#x60;block in read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:144:in &#x60;block in read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:144:in &#x60;block in read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:144:in &#x60;block in read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:144:in &#x60;block in read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:144:in &#x60;block in read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:144:in &#x60;block in read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:144:in &#x60;block in read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:144:in &#x60;block in read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:144:in &#x60;block in read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:144:in &#x60;block in read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:144:in &#x60;block in read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:144:in &#x60;block in read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:144:in &#x60;block in read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:144:in &#x60;block in read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:144:in &#x60;block in read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:144:in &#x60;block in read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:144:in &#x60;block in read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:140:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:119:in &#x60;read&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/s
ite.rb:41:in &#x60;process&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/c
ommand.rb:53:in &#x60;process_site&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/c
ommands/build.rb:50:in &#x60;build&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/c
ommands/build.rb:33:in &#x60;process&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/lib/jekyll/c
ommands/build.rb:17:in &#x60;block (2 levels) in init_with_program&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/mercenary-0.3.4/lib/merce
nary/command.rb:220:in &#x60;call&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/mercenary-0.3.4/lib/merce
nary/command.rb:220:in &#x60;block in execute&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/mercenary-0.3.4/lib/merce
nary/command.rb:220:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/mercenary-0.3.4/lib/merce
nary/command.rb:220:in &#x60;execute&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/mercenary-0.3.4/lib/merce
nary/program.rb:35:in &#x60;go&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/mercenary-0.3.4/lib/merce
nary.rb:22:in &#x60;program&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.2.0/bin/jekyll:1
8:in &#x60;&lt;top (required)&gt;&#x27;
        from C:/tools/ruby200/bin/jekyll:23:in &#x60;load&#x27;
        from C:/tools/ruby200/bin/jekyll:23:in &#x60;&lt;main&gt;&#x27;
&#x60;&#x60;&#x60;
---
> from: [**gep13**](https://github.com/jekyll/jekyll-help/issues/121#issuecomment-52247706) on: **Thursday, August 14, 2014**

@parkr Any ideas on how to proceed with this?  Is there any additional debugging that I can do?  Would you be in a position to jump on the VM via RDP to take a look if I can get you access?
---
> from: [**gep13**](https://github.com/jekyll/jekyll-help/issues/121#issuecomment-52278562) on: **Thursday, August 14, 2014**

I have just installed the latest version of jekyll, 2.3.0, hoping that this might help, but I get the same problem (albeit, slightly different line numbers:

&#x60;&#x60;&#x60;
Configuration file: C:/temp/jekyllrbtest/_config.yml
       Deprecation: The &#x27;pygments&#x27; configuration option has been renamed to &#x27;highlighter&#x27;. Please update your co
e accordingly. The allowed values are &#x27;rouge&#x27;, &#x27;pygments&#x27; or null.
            Source: C:/temp/jekyllrbtest
       Destination: C:/temp/jekyllrbtest/_site
      Generating...
C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:135:in &#x60;open&#x27;: No such file or directo
Errno::ENOENT)
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:135:in &#x60;entries&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:135:in &#x60;block in read_dir
&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:135:in &#x60;chdir&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:135:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:146:in &#x60;block in read_dir
&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:146:in &#x60;block in read_dir
&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:146:in &#x60;block in read_dir
&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:146:in &#x60;block in read_dir
&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:146:in &#x60;block in read_dir
&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:146:in &#x60;block in read_dir
&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:146:in &#x60;block in read_dir
&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:146:in &#x60;block in read_dir
&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:146:in &#x60;block in read_dir
&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:146:in &#x60;block in read_dir
&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:146:in &#x60;block in read_dir
&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:146:in &#x60;block in read_dir
&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:146:in &#x60;block in read_dir
&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:146:in &#x60;block in read_dir
&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:146:in &#x60;block in read_dir
&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:146:in &#x60;block in read_dir
&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:146:in &#x60;block in read_dir
&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:146:in &#x60;block in read_dir
&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:146:in &#x60;block in read_dir
&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:146:in &#x60;block in read_dir
&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:142:in &#x60;read_directories&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:121:in &#x60;read&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/site.rb:43:in &#x60;process&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/command.rb:28:in &#x60;process_site&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/commands/build.rb:55:in &#x60;build&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/commands/build.rb:33:in &#x60;process&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/lib/jekyll/commands/build.rb:17:in &#x60;block (2
 in init_with_program&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/mercenary-0.3.4/lib/mercenary/command.rb:220:in &#x60;call&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/mercenary-0.3.4/lib/mercenary/command.rb:220:in &#x60;block in
&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/mercenary-0.3.4/lib/mercenary/command.rb:220:in &#x60;each&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/mercenary-0.3.4/lib/mercenary/command.rb:220:in &#x60;execute&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/mercenary-0.3.4/lib/mercenary/program.rb:35:in &#x60;go&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/mercenary-0.3.4/lib/mercenary.rb:22:in &#x60;program&#x27;
        from C:/tools/ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.3.0/bin/jekyll:18:in &#x60;&lt;top (required)&gt;&#x27;
        from C:/tools/ruby200/bin/jekyll:23:in &#x60;load&#x27;
        from C:/tools/ruby200/bin/jekyll:23:in &#x60;&lt;main&gt;&#x27;
&#x60;&#x60;&#x60;
---
> from: [**gep13**](https://github.com/jekyll/jekyll-help/issues/121#issuecomment-52359303) on: **Friday, August 15, 2014**

Ok, this is going to be a stoopid question, but here goes...

When I ran the command to scaffold out the site, it created an app folder in the hierarchy.  The above output is for when I run &#x60;&#x60;&#x60;jekyll build --trace&#x60;&#x60;&#x60; on the root folder.  If I cd into the app folder, and run the command again, this is the output I get:

![image](https://cloud.githubusercontent.com/assets/1271146/3939055/f3e6aa0e-24c1-11e4-9be2-918cc00b6b8d.png)

So the stoopid question...

Which folder should I be running the command on?  I suspect that it is the latter, i.e. app, but this still seems to be generating an error in the build, however, if I do &#x60;&#x60;&#x60;jekyll serve --watch&#x60;&#x60;&#x60; I do get a site rendered:

![image](https://cloud.githubusercontent.com/assets/1271146/3939079/2f20cd52-24c2-11e4-8604-2035da9a4ab5.png)

This doesn&#x27;t really explain why the &#x60;&#x60;&#x60;grunt serve&#x60;&#x60;&#x60; command isn&#x27;t working though, or does it?

Any help?
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/121#issuecomment-52382665) on: **Friday, August 15, 2014**

&gt; Which folder should I be running the command on?

Run it at the root of your site. That&#x27;s usually where your &#x60;_config.yml&#x60; file is or where you expect visitors to your site to go to (i.e. has the homepage &#x60;index.html&#x60;).

If &#x60;jekyll serve --watch&#x60; works just fine for you then it&#x27;s something in &#x60;grunt&#x60; that is acting up.
---
> from: [**gep13**](https://github.com/jekyll/jekyll-help/issues/121#issuecomment-52388412) on: **Saturday, August 16, 2014**

Sorry, I am very confused by all of this.

The folder structure that the generator-jekyllrb provided is the following:

Root Folder:
![image](https://cloud.githubusercontent.com/assets/1271146/3941375/68c0c0b8-2525-11e4-9933-471659228d60.png)

App Folder:
![image](https://cloud.githubusercontent.com/assets/1271146/3941376/77d9ef34-2525-11e4-8abd-811a049bb9bd.png)

If I run &#x60;&#x60;&#x60;jekyll build --trace&#x60;&#x60;&#x60; on the root folder, I get the errors mentioned above.  If I run &#x60;&#x60;&#x60;jekyll build --trace&#x60;&#x60;&#x60; on the App Folder, it seems to work.  But the &#x60;&#x60;&#x60;_config.yml&#x60;&#x60;&#x60; and &#x60;&#x60;&#x60;index.html&#x60;&#x60;&#x60; file are not in the same location, therefore, am I right in saying that running jekyll build in the app folder means it isn&#x27;t respecting what is in the config file?  Perhaps @robwierzbowski can help with why they are in different locations.

As you can probably tell, I am very new to this, and am really feeling my way in the dark trying to get things to work.

If this is a problem with grunt, any thoughts on where I should go for some help?

Thanks!
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/121#issuecomment-52389448) on: **Saturday, August 16, 2014**

The Jekyll root directory is the one in the second screenshot.
---
> from: [**gep13**](https://github.com/jekyll/jekyll-help/issues/121#issuecomment-52389493) on: **Saturday, August 16, 2014**

@kleinfreund Thanks for the information.  So that explains why doing &#x60;&#x60;&#x60;jekyll build --trace&#x60;&#x60;&#x60; in the app folder works, and why doing it on the root folder above fails.  Any ideas on why grunt seems to be failing?

Thanks!
---
> from: [**gep13**](https://github.com/jekyll/jekyll-help/issues/121#issuecomment-52389906) on: **Saturday, August 16, 2014**

If I do what it suggests, i.e &#x60;&#x60;&#x60;grunt serve --force&#x60;&#x60;&#x60; it gets further, but the end result is this:

![image](https://cloud.githubusercontent.com/assets/1271146/3941554/bd9ed784-2532-11e4-83c7-b10afb0b5487.png)

---
> from: [**gep13**](https://github.com/jekyll/jekyll-help/issues/121#issuecomment-52839515) on: **Wednesday, August 20, 2014**

Can no-one help with this issue?  I can&#x27;t be the only person who has ran into this.
---
> from: [**AJ-Acevedo**](https://github.com/jekyll/jekyll-help/issues/121#issuecomment-53513088) on: **Tuesday, August 26, 2014**

Lame suggestion but figured I&#x27;d post it.
If you have access to a Mac or Linux machine, you can build your jekyll site there and then host the compiled html on your Windows server.
---
> from: [**gep13**](https://github.com/jekyll/jekyll-help/issues/121#issuecomment-53530552) on: **Tuesday, August 26, 2014**

@AJ-Acevedo thanks for the reply!  While that would work, I guess, I don&#x27;t have a Mac or Linux machine to hand.  Part of the appeal of Jekyll is that I should be able to use it on &quot;any&quot; machine that I want to, but it looks like there is some sort of incompatibility with Windows.

I am actually not looking to host on a Windows Machine, but rather host on Github Pages, but I would like to do the main development work, and compilation on my Windows machine.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/121#issuecomment-55648558) on: **Monday, September 15, 2014**

@gep13 Have you discovered something that will work for you? You might want to try a Vagrant/VirtualBox setup.
---
> from: [**gep13**](https://github.com/jekyll/jekyll-help/issues/121#issuecomment-55650586) on: **Monday, September 15, 2014**

@troyswanson Thank you for your reply! Unfortunately, no, I was never able to find a solution for this :-(  The test that I was doing, listed in the comments above, was on a brand new VM created in Windows Azure.  Am I missing something?  What would using Vagrant/VirtualBox bring to the table?

Thanks!
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/121#issuecomment-55652354) on: **Monday, September 15, 2014**

VirtualBox would provide Vagrant with a platform for running a virtual machine, running Ubuntu or some other Linux distro. You can use that to work with Jekyll much more easily. Here&#x27;s the breakdown if you&#x27;ve never used any of this stuff before:

1. Download/install VirtualBox
2. Download/install Vagrant
3. Use Vagrant to create a new environment
   1. Download a new box template (&#x60;vagrant box add hashicorp/precise64&#x60;)
   2. Initialize a box (&#x60;vagrant init hashicorp/precise64&#x60;)
   3. Fire it up (&#x60;vagrant up&#x60;)
4. SSH into the box you just made (&#x60;vagrant ssh&#x60;)

Now you have a Linux box! There&#x27;s other cool stuff you can do like sharing folders between the host machine and the virtual machine, advanced networking, etc.

Hopefully this helps get you a little bit closer to the holy grail of developing on a *nix system. :wink: 
---
> from: [**gep13**](https://github.com/jekyll/jekyll-help/issues/121#issuecomment-55653650) on: **Monday, September 15, 2014**

Right, sorry, I see what you are getting at :-)

I am familiar with Vagrant/VirtualBox, but to be honest, I am a Windows guy :-D  I develop on Windows all day, mainly in the .Net eco system, and run Windows at home as well.  While I understand that I &quot;could&quot; spin up a VM, and install Linux, I was under the impression that Ruby works happily on both Linux and Windows, so I should be able to work on either, no?

I am not saying that I am against using Linux, I am just trying to understand why I should need to.

I wanted to pick up learning Ruby, by using it as the underlying engine for running my blog, I wasn&#x27;t expecting to have to pick up learning another Operating System, to try it out.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/121#issuecomment-55654602) on: **Monday, September 15, 2014**

Yeah, that&#x27;s fair enough.

Unfortunately, this is the only thing I can offer, since I don&#x27;t have a Windows machine to troubleshoot with! I wish I could help you more, but I just have no way to recreate what&#x27;s happening on your machine so that I can understand the problem better.

I will keep this issue open, though. Hopefully someone with more knowledge of the matter can help out more than I can. :confounded: 
---
> from: [**gep13**](https://github.com/jekyll/jekyll-help/issues/121#issuecomment-55654997) on: **Monday, September 15, 2014**

Not a problem at all, thanks for responding.

In a way, I guess I have &quot;fixed&quot; the problem.  I switched to using Octopress, rather than straight up Jekyll, and so far, that seems to be working really well.  Blog site is up and running, with CI builds happening everytime I publish thanks to AppVeyor.  Quite why Octopress works, and Jekyll doesn&#x27;t, I have no idea.  They were both running on the same machines.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/121#issuecomment-55658838) on: **Monday, September 15, 2014**

Interesting. Well, feel free to close this issue if you&#x27;re good to go. You can leave it open if you want, too.
---
> from: [**gep13**](https://github.com/jekyll/jekyll-help/issues/121#issuecomment-55698707) on: **Monday, September 15, 2014**

Yes, I believe I am all set.  It&#x27;s a shame that we couldn&#x27;t get to the bottom of this one though, as I suspect it might put people off getting started with Jekyll. :-(
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/121#issuecomment-55700817) on: **Monday, September 15, 2014**

It looks like you weren&#x27;t using straight-up Jekyll after all – by running &#x60;grunt serve&#x60;, you were using a different system that wrapped around Jekyll. It&#x27;s not difficult to imagine that that is where things could have been going wrong.

Your error (the duplicate drive) is reminiscent of a bug we found in v1.5.3 of Jekyll, but it was patched by the time v2.0 came out. I think you may have been running v1.5.3, because the last line of a failed build should show jekyll and its version, not &#x60;error:&#x60;. Did you check the version that &#x60;grunt&#x60; was using? It looks like generator-jekyllrb was out of date.

The warning about posix-spawn is fine, but you can set &#x60;highlighter: rouge&#x60; in Jekyll 2.0.
---
> from: [**gep13**](https://github.com/jekyll/jekyll-help/issues/121#issuecomment-55701679) on: **Monday, September 15, 2014**

That is actually a fair point, I had actually forgotten that I was trying to use the generator!  I had tried various things while trying to get this all setup.  For now, I am very happy with the Octopress setup that I have running, so I will continue with that.

Thanks for all the help!
