# [Jekyll not compiling site on localhost:8000](https://github.com/jekyll/jekyll-help/issues/184)

> state: **closed** opened by: **** on: ****

Hello,

I have Jekyll 2.4.0 and Ruby 2.0.09576 installed.

I am working along with Travis from YouTube channel &#x27;DevTips&#x27; and using all of his files/information to compile his &#x27;Artist&#x27; project site. Last week I successfully served the project site via Ruby/Jekyll serve from source folder. Today, I tried the same process and the site would not compile. I am using same localhost:8000, that is verified. I am running Win7 64bit and followed advice to run UTF-8 encoding command: chcp 65001.

No files have changed in the portfolio folder on my PC. The other thing on jekyllrb.com/windows/ says is to add a line of code for &#x27;Auto-regeneration&#x27; to the &#x27;Gemfile&#x27;...OK, it does not say where this &#x27;Gemfile&#x27; is to edit...is it here?: C:.../RubyDevKit/bin/gem.windows batch file?

I am stumped as to why the project files compiled and displayed just fine last week and today it does not work, NO connection can be made to localhost:8000...so no site generated. I also tried to view view site via Prepros, using their &#x27;Open Live Preview&#x27; and that generates a site on localhost:8001 displaying only this text: --- --- {% include header.html %} {% include about.html %} {% include work.html %} {% include clients.html %} {% include contact.html %} {% include form.html %} {% include footer.html %} 

Any ideas?

Thank you!
Mark



### Comments

---
> from: [**mloderme**](https://github.com/jekyll/jekyll-help/issues/184#issuecomment-61837163) on: **Wednesday, November 05, 2014**

Anybody out there? Anybody home? NO..OK...sweet jeebuz I am all alone, what a feeling!
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/184#issuecomment-61844491) on: **Wednesday, November 05, 2014**

When you run &#x60;jekyll build&#x60; from the source folder, what output do you receive? I&#x27;m curious to know why you&#x27;re using localhost:8000, also; the default port for Jekyll is 4000.

Gemfiles are a thing from Bundler. Learn more about them [here](http://bundler.io/gemfile.html).
---
> from: [**mloderme**](https://github.com/jekyll/jekyll-help/issues/184#issuecomment-61856919) on: **Wednesday, November 05, 2014**

Hello,

              

I am new to this stuff, have been following Travis, author of YouTube channel &#x27;DevTips&#x27; and am using his files. To be clear, I successfully served the site on localhost:8000 last week. I&#x27;ve also configured the _config.yml file to ensure port used is 8000, but also deleted that and used default port 4000, to no avail.

Thank you!

Mark

Date: Wed, 5 Nov 2014 09:11:56 -0800
From: notifications@github.com
To: jekyll-help@noreply.github.com
CC: mloderme@hotmail.com
Subject: Re: [jekyll-help] Jekyll not compiling site on localhost:8000 (#184)

When you run jekyll build from the source folder, what output do you receive? I&#x27;m curious to know why you&#x27;re using localhost:8000, also; the default port for Jekyll is 4000.


Gemfiles are a thing from Bundler. Learn more about them here.


—
Reply to this email directly or view it on GitHub.

  
    
    
  
  
 		 	   		  =
---
> from: [**mloderme**](https://github.com/jekyll/jekyll-help/issues/184#issuecomment-61865895) on: **Wednesday, November 05, 2014**

Hello,

New jekyll output with error after reoading all source files from https://github.com/DevTips/Artists-Theme then, restarting Ruby and running: jekyll serve --watch from same source folder.

              

Date: Wed, 5 Nov 2014 09:11:56 -0800
From: notifications@github.com
To: jekyll-help@noreply.github.com
CC: mloderme@hotmail.com
Subject: Re: [jekyll-help] Jekyll not compiling site on localhost:8000 (#184)

When you run jekyll build from the source folder, what output do you receive? I&#x27;m curious to know why you&#x27;re using localhost:8000, also; the default port for Jekyll is 4000.


Gemfiles are a thing from Bundler. Learn more about them here.


—
Reply to this email directly or view it on GitHub.

  
    
    
  
  
 		 	   		  =
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/184#issuecomment-63156318) on: **Friday, November 14, 2014**

You may also use &#x60;bundle exec jekyll serve --port 8000&#x60;.
