---
layout: page
title: Resources
---

Here is a list of resources that you may find helpful throughout the semester.  

------------------

## Textbooks

- Distributed Systems: Principles and Paradigms, Andrew Tanenbaum and Maarten van Steen, Prentice Hall.
- Principles of Computer System Design. Jerome Saltzer and M. Frans Kaashoek, Morgan Kaufmann. 
- Advanced Programming in the UNIX Environment. W. Richard Stevens. 

------------------

## Git

- [Atlassian's Git tutorial](https://www.atlassian.com/git/tutorials/what-is-version-control) A fairly comprehensive tutorial that doesn't make your head hurt.

------------------

## Go

- [Tour of Go](https://tour.golang.org/list) An interactive tool for exploring basic Go code.  Good for getting a feel for the syntax.
- [Go Documentation](https://golang.org/doc/) Good for reference.
- \*[Concurrency in Go](https://golang.org/doc/effective_go.html#concurrency) Information on Channels and goroutines
- \*[Go RPC Documentation](https://golang.org/pkg/net/rpc/) Go's Remote Procedure Call library documentation

\*<sub>You will be returning to these frequently for assignments 2-4!</sub>

------------------

## Operating Systems

- Operating Systems: Principles and Practice, Thomas Anderson and Mike Dahlin,
  second edition recommended.
- Modern Operating Systems, Andrew Tanenbaum, third edition recommended.

------------------

## Research Papers

- [MapReduce Paper](https://static.googleusercontent.com/media/research.google.com/en//archive/mapreduce-osdi04.pdf) Google's MapReduce system for distributed computation.  This work inspired the creation of the open source version, Apache Spark (developed at Berkeley). Familiarity with this paper, particularly Section 3.3: Fault Tolerance, is necessary for assignment 1.
- [Time, Clocks, and the Ordering of Events in a Distributed System](http://lamport.azurewebsites.net/pubs/time-clocks.pdf) A foundational distributed systems paper that discusses the difficulties of ordering events, and introduces the concepts of "Lamport Clocks" and "Vector Clocks".
- [Linearizability](https://cs.brown.edu/~mph/HerlihyW90/p463-herlihy.pdf)
  Another foundational distributed systems paper written by Jeannette Wing, the
  Director of the Data Science Institute at Columbia.  Focuses on Linearizability, a strong consistency model for distributed data storage.
- [CAP Theorem](https://users.ece.cmu.edu/~adrian/731-sp04/readings/GL-cap.pdf) A very easy-to-follow impossibility proof showing that no distributed system can all three of the properties: "Consistency, Availability, and Partition tolerance."
- [Part-Time Parliament](http://lamport.azurewebsites.net/pubs/lamport-paxos.pdf) A bit of history: Leslie Lamport's first paper on the Paxos protocol.  Reads like a story more than a research paper. Virtually no one understood this paper, which led to author publishing:
- [Paxos Made Simple](http://lamport.azurewebsites.net/pubs/paxos-simple.pdf) Which is what you will be implementing in assignment 3.

------------------
