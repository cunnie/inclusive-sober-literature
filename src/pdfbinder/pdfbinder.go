package pdfbinder

func Bind() string {
	// open the files
	// for each file
	  // paginates file
		// Paginator
		// gets a file
		// for each "PAGE 0X"
			// writes text in that page to 0X.md
	// makes the mark-up
		// gets a list of files (pages, 0X.md)
		// for each missing 0Y.md file < initial 0X.md
			// generate empty <p class="page-break-after"></p>
		// for each file 0X.md
			// generate <p class="page-break-after">{{  toHTML .DirPath "0X.md"  }}</p>
		// puts output of for loops into html body
			// <!doctype html>
			// <html lang="en">
			//  <head>
			//    <meta charset="utf-8">
			//    <title>My PDF</title>
			//  </head>
			//  <body>
			//  	{{ body, aka all dem P tags }}
			//  </body>
			// </html>
		// saves .html
	// makes the right gotenberg call using gotenberg go library
	// returns pdf file
	return "yellow"
}


// func TheBestFunction(io.File[] inputFiles) (io.File, io.File[]