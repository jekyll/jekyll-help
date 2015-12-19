# [Jekyll (Kramdown) removes closing PRE tag](https://github.com/jekyll/jekyll-help/issues/236)

> state: **closed** opened by: **** on: ****

I have this post:

	---
	layout: post
	status: publish
	published: true
	title: Handling keypresses in Cocoa games
	author:
	  display_name: uliwitness
	  login: uliwitness
	  email: xxx@example.com
	  url: &#x27;&#x27;
	author_login: uliwitness
	author_email: xxx@example.com
	wordpress_id: 736
	wordpress_url: http://orangejuiceliberationfront.com/?p=736
	date: &#x27;2014-12-13 11:57:08 +0100&#x27;
	date_gmt: &#x27;2014-12-13 11:57:08 +0100&#x27;
	categories:
	- Macintosh
	- Objective-C
	- Programming
	- Game Design
	tags: []
	comments: []
	---
	&lt;p&gt;The solution is to keep track of which key is down yourself. Override &lt;tt&gt;-keyDown&lt;/tt&gt; and &lt;tt&gt;-keyUp&lt;/tt&gt; to keep track of which keys are being held down. I&#x27;m using a C++ &lt;tt&gt;unordered_set&lt;/tt&gt; for that, but an Objective-C &lt;tt&gt;NSIndexSet&lt;/tt&gt; would work just as well:&lt;/p&gt;
	&lt;pre&gt;
	@interface ICGMapView : NSView
	{
		std::unordered_set&lt;unichar&gt;	pressedKeys;
	}
	
	@end
	&lt;/pre&gt;
	&lt;p&gt;and in the implementation:&lt;/p&gt;
	&lt;pre&gt;
	-(void)	keyDown:(NSEvent *)theEvent
	{
		NSString	*	pressedKeyString = theEvent.charactersIgnoringModifiers;
		unichar			pressedKey = (pressedKeyString.length &gt; 0) ? [pressedKeyString characterAtIndex: 0] : 0;
		if( pressedKey )
			pressedKeys.insert( pressedKey );
	}
	
	
	-(void)	keyUp:(NSEvent *)theEvent
	{
		NSString	*	pressedKeyString = theEvent.charactersIgnoringModifiers;
		unichar			pressedKey = (pressedKeyString.length &gt; 0) ? [pressedKeyString characterAtIndex: 0] : 0;
		if( pressedKey )
		{
			auto foundKey = pressedKeys.find( pressedKey );
			if( foundKey != pressedKeys.end() )
				pressedKeys.erase(foundKey);
		}
	}
	&lt;/pre&gt;

And when I &#x60;jekyll build&#x60; this file, it removes the first &amp;lt;/pre&amp;gt; tag (after &quot;@end&quot;):

	&lt;!DOCTYPE html&gt;
	&lt;html&gt;
	
	  &lt;head&gt;
	  &lt;meta charset=&quot;utf-8&quot;&gt;
	  &lt;meta http-equiv=&quot;X-UA-Compatible&quot; content=&quot;IE=edge&quot;&gt;
	  &lt;meta name=&quot;viewport&quot; content=&quot;width=device-width, initial-scale=1&quot;&gt;
	
	  &lt;title&gt;Handling keypresses in Cocoa games&lt;/title&gt;
	  &lt;meta name=&quot;description&quot; content=&quot;The solution is to keep track of which key is down yourself. Override -keyDown and -keyUp to keep track of which keys are being held down. I&#x27;m using a C++ un...&quot;&gt;
	
	  &lt;link rel=&quot;stylesheet&quot; href=&quot;/css/main.css&quot;&gt;
	  &lt;link rel=&quot;canonical&quot; href=&quot;http://orangejuiceliberationfront.com/macintosh/objective-c/programming/game%20design/2014/12/13/handling-keypresses-in-cocoa-games.html&quot;&gt;
	  &lt;link rel=&quot;alternate&quot; type=&quot;application/rss+xml&quot; title=&quot;Orange Juice Liberation Front&quot; href=&quot;http://orangejuiceliberationfront.com/feed.xml&quot; /&gt;
	&lt;/head&gt;
	
	
	  &lt;body&gt;
	
		&lt;header class=&quot;site-header&quot;&gt;
	
	  &lt;div class=&quot;wrapper&quot;&gt;
	
		&lt;a class=&quot;site-title&quot; href=&quot;/&quot;&gt;Orange Juice Liberation Front&lt;/a&gt;
	
		&lt;nav class=&quot;site-nav&quot;&gt;
		  &lt;a href=&quot;#&quot; class=&quot;menu-icon&quot;&gt;
			&lt;svg viewBox=&quot;0 0 18 15&quot;&gt;
			  &lt;path fill=&quot;#424242&quot; d=&quot;M18,1.484c0,0.82-0.665,1.484-1.484,1.484H1.484C0.665,2.969,0,2.304,0,1.484l0,0C0,0.665,0.665,0,1.484,0 h15.031C17.335,0,18,0.665,18,1.484L18,1.484z&quot;/&gt;
			  &lt;path fill=&quot;#424242&quot; d=&quot;M18,7.516C18,8.335,17.335,9,16.516,9H1.484C0.665,9,0,8.335,0,7.516l0,0c0-0.82,0.665-1.484,1.484-1.484 h15.031C17.335,6.031,18,6.696,18,7.516L18,7.516z&quot;/&gt;
			  &lt;path fill=&quot;#424242&quot; d=&quot;M18,13.516C18,14.335,17.335,15,16.516,15H1.484C0.665,15,0,14.335,0,13.516l0,0 c0-0.82,0.665-1.484,1.484-1.484h15.031C17.335,12.031,18,12.696,18,13.516L18,13.516z&quot;/&gt;
			&lt;/svg&gt;
		  &lt;/a&gt;
	
		  &lt;div class=&quot;trigger&quot;&gt;
			
		  &lt;/div&gt;
		&lt;/nav&gt;
	
	  &lt;/div&gt;
	
	&lt;/header&gt;
	
	
		&lt;div class=&quot;page-content&quot;&gt;
		  &lt;div class=&quot;wrapper&quot;&gt;
			&lt;div class=&quot;post&quot;&gt;
	
	  &lt;header class=&quot;post-header&quot;&gt;
		&lt;h1 class=&quot;post-title&quot;&gt;Handling keypresses in Cocoa games&lt;/h1&gt;
		&lt;p class=&quot;post-meta&quot;&gt;Dec 13, 2014 • {&quot;display_name&quot;=&gt;&quot;uliwitness&quot;, &quot;login&quot;=&gt;&quot;uliwitness&quot;, &quot;email&quot;=&gt;&quot;witness.of.teachtext@gmx.net&quot;, &quot;url&quot;=&gt;&quot;&quot;}&lt;/p&gt;
	  &lt;/header&gt;
	
	  &lt;article class=&quot;post-content&quot;&gt;
		&lt;p&gt;The solution is to keep track of which key is down yourself. Override &lt;tt&gt;-keyDown&lt;/tt&gt; and &lt;tt&gt;-keyUp&lt;/tt&gt; to keep track of which keys are being held down. I&#x27;m using a C++ &lt;tt&gt;unordered_set&lt;/tt&gt; for that, but an Objective-C &lt;tt&gt;NSIndexSet&lt;/tt&gt; would work just as well:&lt;/p&gt;
	&lt;pre&gt;
	@interface ICGMapView : NSView
	{
		std::unordered_set&lt;unichar&gt;	pressedKeys;
	}
	
	@end
	
	&lt;p&gt;and in the implementation:&lt;/p&gt;
	&lt;pre&gt;
	-(void)	keyDown:(NSEvent *)theEvent
	{
		NSString	*	pressedKeyString = theEvent.charactersIgnoringModifiers;
		unichar			pressedKey = (pressedKeyString.length &amp;gt; 0) ? [pressedKeyString characterAtIndex: 0] : 0;
		if( pressedKey )
			pressedKeys.insert( pressedKey );
	}
	
	
	-(void)	keyUp:(NSEvent *)theEvent
	{
		NSString	*	pressedKeyString = theEvent.charactersIgnoringModifiers;
		unichar			pressedKey = (pressedKeyString.length &amp;gt; 0) ? [pressedKeyString characterAtIndex: 0] : 0;
		if( pressedKey )
		{
			auto foundKey = pressedKeys.find( pressedKey );
			if( foundKey != pressedKeys.end() )
				pressedKeys.erase(foundKey);
		}
	}
	&lt;/pre&gt;
	&lt;/unichar&gt;&lt;/pre&gt;
	
	  &lt;/article&gt;
	
	&lt;/div&gt;
	
		  &lt;/div&gt;
		&lt;/div&gt;
	
		&lt;footer class=&quot;site-footer&quot;&gt;
	
	  &lt;div class=&quot;wrapper&quot;&gt;
	
		&lt;h2 class=&quot;footer-heading&quot;&gt;Orange Juice Liberation Front&lt;/h2&gt;
	
		&lt;div class=&quot;footer-col-wrapper&quot;&gt;
		  &lt;div class=&quot;footer-col  footer-col-1&quot;&gt;
			&lt;ul class=&quot;contact-list&quot;&gt;
			  &lt;li&gt;Orange Juice Liberation Front&lt;/li&gt;
			  &lt;li&gt;&lt;a href=&quot;mailto:&quot;&gt;&lt;/a&gt;&lt;/li&gt;
			&lt;/ul&gt;
		  &lt;/div&gt;
	
		  &lt;div class=&quot;footer-col  footer-col-2&quot;&gt;
			&lt;ul class=&quot;social-media-list&quot;&gt;
			  
	
			  
			  &lt;li&gt;
				&lt;a href=&quot;https://twitter.com/uliwitness&quot;&gt;
				  &lt;span class=&quot;icon  icon--twitter&quot;&gt;
					&lt;svg viewBox=&quot;0 0 16 16&quot;&gt;
					  &lt;path fill=&quot;#828282&quot; d=&quot;M15.969,3.058c-0.586,0.26-1.217,0.436-1.878,0.515c0.675-0.405,1.194-1.045,1.438-1.809
					  c-0.632,0.375-1.332,0.647-2.076,0.793c-0.596-0.636-1.446-1.033-2.387-1.033c-1.806,0-3.27,1.464-3.27,3.27 c0,0.256,0.029,0.506,0.085,0.745C5.163,5.404,2.753,4.102,1.14,2.124C0.859,2.607,0.698,3.168,0.698,3.767 c0,1.134,0.577,2.135,1.455,2.722C1.616,6.472,1.112,6.325,0.671,6.08c0,0.014,0,0.027,0,0.041c0,1.584,1.127,2.906,2.623,3.206 C3.02,9.402,2.731,9.442,2.433,9.442c-0.211,0-0.416-0.021-0.615-0.059c0.416,1.299,1.624,2.245,3.055,2.271 c-1.119,0.877-2.529,1.4-4.061,1.4c-0.264,0-0.524-0.015-0.78-0.046c1.447,0.928,3.166,1.469,5.013,1.469 c6.015,0,9.304-4.983,9.304-9.304c0-0.142-0.003-0.283-0.009-0.423C14.976,4.29,15.531,3.714,15.969,3.058z&quot;/&gt;
					&lt;/svg&gt;
				  &lt;/span&gt;
	
				  &lt;span class=&quot;username&quot;&gt;uliwitness&lt;/span&gt;
				&lt;/a&gt;
			  &lt;/li&gt;
			  
			&lt;/ul&gt;
		  &lt;/div&gt;
	
		  &lt;div class=&quot;footer-col  footer-col-3&quot;&gt;
			&lt;p class=&quot;text&quot;&gt;&lt;/p&gt;
		  &lt;/div&gt;
		&lt;/div&gt;
	
	  &lt;/div&gt;
	
	&lt;/footer&gt;
	
	
	  &lt;/body&gt;
	
	&lt;/html&gt;

Let me know if you need more info or the original files.

NB -- there were oodles of empty lines in the trigger div which I deleted for this report. The original post was generated by &#x60;jekyll-import&#x60; from a WordPress post, for which I can also provide the original if needed.

### Comments

---
> from: [**uliwitness**](https://github.com/jekyll/jekyll-help/issues/236#issuecomment-69706110) on: **Monday, January 12, 2015**

Nevermind, found the issue: It&#x27;s the &quot;&amp;lt;unichar&amp;gt;&quot; in the example code. Changing that to &amp;amp;lt;unichar&amp;amp;gt; on the Wordpress server fixed the removal of the Pre tag.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/236#issuecomment-69820189) on: **Tuesday, January 13, 2015**

Hey @uliwitness, I would highly encourage you to use code highlighting in your Markdown files instead of including straight HTML in them.

You&#x27;d be able to get away with much more from an input perspective, plus your code would render with fancy colors!
---
> from: [**uliwitness**](https://github.com/jekyll/jekyll-help/issues/236#issuecomment-71071646) on: **Thursday, January 22, 2015**

This is a WordPress import, so I might eventually do that, but right now I just want to migrate the blog intact :-)
