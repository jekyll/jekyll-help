# [Jekyll&#x27;s installation process shenanigans.](https://github.com/jekyll/jekyll-help/issues/245)

> state: **open** opened by: **** on: ****

Greetings fellas,

Been struggilng with the installation process for windows(i can&#x27;t afford a mac), following some very precise tutorials i got myself around installing ruby and a devkit(both are compatible with each other).

But when i try to run the jekyll cmd on the console(not windows&#x27;s) i get an error message as if there was an error while fetching data from jekyll&#x27;s server.

Any thought/recommendations are welcome, thanks for your time in advance and have a great day!.

- Ysag

### Comments

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/245#issuecomment-71097679) on: **Thursday, January 22, 2015**

That error message you were talking about, where is it? That&#x27;d be super helpful to help you out.
---
> from: [**Ysag**](https://github.com/jekyll/jekyll-help/issues/245#issuecomment-71098833) on: **Thursday, January 22, 2015**

ERROR: Could not find a valid gem &#x27;jekyll&#x27; &lt;&gt;=0&gt;, here is why:
	Unable to download data from https://rubygems.org/ - SSL_connect returned=1
	errno=0 state=SSLv3 read server certificate B: certificate verify failed &lt;https://api.rubygems.org/latest_specs.4.8.gz&gt;

Here&#x27;s the exact same message i am getting. Thanks for your help btw.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/245#issuecomment-71099127) on: **Thursday, January 22, 2015**

Check out this [SO answer](http://stackoverflow.com/a/27641786).
---
> from: [**Ysag**](https://github.com/jekyll/jekyll-help/issues/245#issuecomment-71102658) on: **Thursday, January 22, 2015**

Thanks for the info Travis!. Though it tells me to run the setup.rb, I assumed to run it with ruby but it is not picking it up. Do i have to have these files on the same dir as for ruby?.

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/245#issuecomment-71103060) on: **Thursday, January 22, 2015**

I don&#x27;t know 100% but I wouldn&#x27;t expect so. You should just be able to unzip the file into some folder somewhere, and then run &#x60;ruby setup.rb&#x60; while you&#x27;re in that folder.

PS: I&#x27;m Troy. :smile: 
---
> from: [**Ysag**](https://github.com/jekyll/jekyll-help/issues/245#issuecomment-71103631) on: **Thursday, January 22, 2015**

Sorry for that namechange all of a sudden!. The thing is that it&#x27;s not letting me run it, it&#x27;s not picking the file type(extension) and i tried going to the folder and running it but it won&#x27;t. 

Sorry for taking all of your time people!, but i just of now got started with ruby and jekyll. I&#x27;ve only used ST for offline Front-end designing.

-Ysag
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/245#issuecomment-71109912) on: **Thursday, January 22, 2015**

Oh, try being inside a &#x60;cmd&#x60; window, &#x60;cd&#x60; into that folder, and try running &#x60;ruby setup.rb&#x60; there.
