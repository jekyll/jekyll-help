# [generating global variable on build that is available to all Liquid::Tags](https://github.com/jekyll/jekyll-help/issues/133)

> state: **closed** opened by: **** on: ****

I&#x27;m writing a liquid tag that needs to check for a file&#x27;s existence in an AWS S3 bucket. Right now I have it making a request to S3 every time the tag occurs. Looks something like:

&#x60;&#x60;&#x60;ruby
   class RenderTag &lt; Liquid::Tag
   ...
    def render(context)
      s3_API.find(filename) #makes a call to S3 API
    end
&#x60;&#x60;&#x60;
This works, but it&#x27;s slow.

It&#x27;d be a lot easier for me to generate a globally accessible object one time and with a single call to S3&#x27;s API prior to site generation containing a list of the files in the s3 bucket. This way, instead of each occurrence of my tag making its own call to S3, it could simply check the already generated directory object. 

&#x60;&#x60;&#x60;ruby
   class RenderTag &lt; Liquid::Tag
   ...
    def render(context)
      #doesn&#x27;t make a call to API just checks the object generated at the beginning of build
      globalDirObject.find(filename) 
    end
&#x60;&#x60;&#x60;

I don&#x27;t know if something like this is even close to possible, but if it is I&#x27;d love some advice. Thanks!

### Comments

