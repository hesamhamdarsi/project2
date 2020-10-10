//encode and decode JSON is like Marshal and unmarshal
//we use Writer to encode and Reader to decode
//writer could be for example a file, or a web connection to client. what ever it is you can encode your
//data to JSON when it wants to go to that writer
//for encode, you first give a writer (wirte in to that output source) and then you encode it
//for decode is the same:
//reader could be for example a file, or a web connection from client. what ever it is you can decode your
//data from JSON when it wants to come from that reader
//for decode, you first give a reader (wirte in to that output source) and then you decode it

//what is exactly a writer:
//it is a type (type writer) of interface