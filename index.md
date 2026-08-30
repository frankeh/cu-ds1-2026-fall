---
layout: default
title: Overview
---

**Time** : Fridays 10:10am-12:40pm <br/>
**Class Location**: CSB 451<br/>
**ED**: Use courseworks link<br/>
**Office Hours**: Posted in [Staff and Office Hours]({{ site.baseurl
}}/04-staff)

------------------

Distributed systems help programmers aggregate the resources of many networked computers to construct highly available and scalable services. Most of the applications and services we interact with today are distributed, some at enormous scales.

This class teaches design and implementation techniques that enable the building of fast, scalable, fault-tolerant distributed systems.
Topics include distributed communication models (e.g., sockets, remote procedure calls, distributed shared memory), distributed synchronization (clock synchronization, logical clocks, distributed mutex), distributed file systems, replication, consistency models, fault tolerance, distributed transactions, agreement and commitment, Paxos-based consensus, MapReduce infrastructures, scalable distributed databases.

The class combines concepts and algorithms with descriptions of real-world implementations at Google, Facebook, Yahoo, Microsoft, LinkedIn, etc.
In addition to lectures, students will get hands-on experience  building distributed systems through a series of coding-oriented homeworks.
The series, adopted from MIT's course, implements a fault-tolerant, sharded key/value store.

------------------

## Grading

The grade will be assigned on performance of five [homeworks]({{ site.baseurl
}}/02-homeworks) and a final exam.
Additionally, a 10% extra credit may be awarded to students with significant and
particularly insightful contributions on ED and/or in class throughout the semester.   There is no specific number of these awards, but you should think of awardees as people who have stood out consistently and have improved the class in some significant way.  Finally, a 10% extra credit is given to students who complete and score perfectly on all eight homework components.
The [first lecture]({{ site.baseurl }}/01-lectures/) details the grading procedure, and the [Deadlines]({{ site.baseurl }}/03-deadlines/) page specifies the deadlines/dates of each homework/quiz.

------------------

## Prerequisites

The homework series will require a lot of coding.  Hence, in this class, we require that you have solid coding experience, particularly
building systems-level components (e.g., not just apps).
This can come either from personal or industry experience, or from the following Columbia courses or equivalents:

1. COMS W3137 Data Structures and Algorithms
2. COMS W3157 Advanced Programming
3. COMS W3827 Fundamentals of Computer Systems
4. W4118 Operating Systems is not required, but it is a big plus for your homework assignments

Please make sure you can meet the resource requirements listed in the  [homeworks]({{ site.baseurl }}/02-homeworks) section.

------------------

## Acknowledgements

This class, along with the materials distributed for it, was inspired by Distributed Systems courses at various institutions:

### Lectures:
  * University of Washington's distributed systems graduate course, [Steve Gribble's version](http://courses.cs.washington.edu/courses/csep552/13sp/);
  * New York University's distributed systems course, [Jinyang Li's version](http://news.cs.nyu.edu/~jinyang/fa10/);
  * CMU's distributed systems course, [Dave Andersen's version](http://www.cs.cmu.edu/~dga/15-440/F10/);
  * University of Southern California's distributed systems class, [Wyatt Lloyd's version](http://www.cs.princeton.edu/~wlloyd/classes/657f16/index.html).

### Homeworks:
  * Homeworks 1-4 use MIT's Go assignment series, [Spring 2015 edition](http://nil.csail.mit.edu/6.824/2015/)\*.
  * Homework 5 is developed by us, in continuation to the MIT series and inspired by University of Washington's [Distributed Systems Labs and Framework](https://github.com/emichael/dslabs).
  
