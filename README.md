GO language immutable structure-sharing collection classes
==========================================================

<!--
	Copyright 2009 The Go Authors. All rights reserved.
	Use of this source code is governed by a BSD-style
	license that can be found in the LICENSE file.
-->

	
		<div id="short-nav">
			<dl>
			<dd><code>import "github.com/eobrain/immut"</code></dd>
			</dl>
			<dl>
			<dd><a href="#overview" class="overviewLink">Overview</a></dd>
			<dd><a href="#index">Index</a></dd>
			
			
			</dl>
		</div>
		<!-- The package's Name is printed as title by the top-level template -->
		<div id="overview" class="toggleVisible">
			<div class="collapsed">
				<h2 class="toggleButton" title="Click to show Overview section">Overview ▹</h2>
			</div>
			<div class="expanded">
				<h2 class="toggleButton" title="Click to hide Overview section">Overview ▾</h2>
				<p>
Immutable Structure-Sharing Types
</p>

			</div>
		</div>
		
	
		<h2 id="index">Index</h2>
		<!-- Table of contents for API; must be named manual-nav to turn off auto nav. -->
		<div id="manual-nav">
			<dl>
			
			
			
			
				
				<dd><a href="#Item">type Item</a></dd>
				
				
			
				
				<dd><a href="#Seq">type Seq</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#List">func List(item ...Item) Seq</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#Remove">func Remove(xs Seq, x Item) Seq</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#Set">func Set(item ...Item) Seq</a></dd>
				
				
			
			
		</dl>

		

		
			<h4>Package files</h4>
			<p>
			<span style="font-size:90%">
			
				<a href="/target/immut.go">immut.go</a>
			
				<a href="/target/list.go">list.go</a>
			
				<a href="/target/null.go">null.go</a>
			
				<a href="/target/set.go">set.go</a>
			
			</span>
			</p>
		
	
		
		
		
		
			
			
			<h2 id="Item">type <a href="/target/immut.go?s=673:694#L9">Item</a></h2>
			<pre>type Item interface{}</pre>
			<p>
An item in the seq
</p>


			

			

			

			

			
		
			
			
			<h2 id="Seq">type <a href="/target/immut.go?s=804:2035#L13">Seq</a></h2>
			<pre>type Seq interface {
    <span class="comment">//O(n) return number of elements</span>
    Len() int
    <span class="comment">//O(n) or O(log(n)) whether item is in seq</span>
    Contains(Item) bool
    <span class="comment">//O(1) or O(log(n)) return first item, or an error if seq is empty</span>
    Front() (Item, error)
    <span class="comment">//O(1) or O(???) return new list with all except the first item</span>
    <span class="comment">//or an error if seq is empty</span>
    Rest() (Seq, error)
    <span class="comment">//O(1) is this the empty seq</span>
    IsEmpty() bool
    <span class="comment">//O(n) Apply the function to each item in the seq</span>
    Each(func(Item))
    <span class="comment">//O(???) Return a concatentaion of the string representations of the items separated by sep</span>
    Join(sep string) string
    <span class="comment">//O(n) or O(???) return a new seq with the item added on to the end</span>
    Add(Item) Seq
    <span class="comment">//return a new seq that is a concatenation of this seq with the given one</span>
    AddAll(Seq) Seq

    <span class="comment">//whether function is true for all items, or if there are no items</span>
    Forall(func(Item) bool) bool

    <span class="comment">//return a new seq where each item is the result of running the function on the corresponding item of this seq</span>
    Map(func(Item) Item) Seq
    <span class="comment">//return a new seq with a subset of the items for which the function is true</span>
    Filter(func(Item) bool) Seq
    <span class="comment">// contains filtered or unexported methods</span>
}</pre>
			<p>
An immutable sequence of Items
Where multiple O(...) given, first is for list, second is for tree set
</p>


			

			

			

			
				
				<h3 id="List">func <a href="/target/list.go?s=674:701#L10">List</a></h3>
				<pre>func List(item ...Item) Seq</pre>
				<p>
Create a new list containing the arguments
</p>

				
			
				
				<h3 id="Remove">func <a href="/target/immut.go?s=2144:2175#L50">Remove</a></h3>
				<pre>func Remove(xs Seq, x Item) Seq</pre>
				<p>
Return sequence resulting from removing the item, or the sequence
itself if item not contained in it
</p>

				
			
				
				<h3 id="Set">func <a href="/target/set.go?s=954:980#L24">Set</a></h3>
				<pre>func Set(item ...Item) Seq</pre>
				<p>
Create a new set containing the arguments
</p>

				
			

			
		
		</div>
	

	







