---
layout: page
title: Homeworks
---

This course consists of a series of five heavy programming assignments.
The first four assignments are based on those used in MIT's distributed systems
course (6.824) while the last assignment is developed at Columbia to complement an
important missing component.
All assignments are in the [Go programming language](https://www.golang.org).
The series builds in stages a simple but fairly realistic, fault-tolerant,
consistent, distributed key-value store, and then model checks a key protocol
for it.

Some of the assignments build off of each other, so it is **critical** that you
begin them early, and have working versions of each other by the deadlines.
The lectures cover the core concepts and protocols underlying the homeworks,
before those homeworks are due, but not always when the homework is first
released.  However, we do provide the necessary materials ahead of time and we
expect you to read them and begin working on the homework _before_ the topic is
fully covered in class.

Instructions for each homework are available within the git repository that will be set up for you following the first day of lecture (as part of HW0). A link will be distributed to you to bootstrap HW0.

Deadlines are specified in each homework's instructions, summarized in the
[Deadlines page]({{ site.baseurl }}/03-deadlines/), and available in the class calendar.

There will be **NO DEADLINE EXTENSIONS** for any reason other than health conditions
(of the student or of someone close if that impacts the student). However, there is a
**72-hour grace period**, accumulated over all homeworks, for which you will not be
downgraded.  Lateness is accounted at hour granularity (i.e., 1 second late == 1 hour late).
Once you reach 72 hours of lateness, the next homework to incur even a one-second delay
will be graded as **ZERO**. Thus, the strong recommendation is to submit on time, even
if with an only somewhat working homework.

## Overview of the series:

- **Assignment 0: Collaboration Policy** (not graded, but required)  
  - Read the class policies and resources, and set up your GitHub account in the class' GitHub Classroom.  
  - Gets students bootstrapped into the course.  

- **Assignment 1: MapReduce**  
  - Build a simple MapReduce library as a way to learn Go.  
  - Introduces the most basic form of fault tolerance in distributed systems: stateless servers.  

- **Assignment 2: Primary/Backup Server**  
  - Develop a simple key/value server that achieves fault tolerance using an in-house protocol and the primary/backup architecture.  
  - Serves as a first step toward understanding the challenges of making stateful servers fault tolerant.  

- **Assignment 3: Paxos**  
  - Implement a fault-tolerant key/value store based on the Paxos protocol.  
  - Provides in-depth understanding of the protocol's intricacies in practice, even in a simplified setting.  

- **Assignment 4: Sharded Key/Value Server**  
  - Extend the Paxos-based store from Assignment 3 by sharding it across multiple replica groups for scalability.  
  - Teaches a common architecture for building fault-tolerant, scalable stateful services, used in most modern storage systems (e.g., Spanner, studied in class).  

- **Assignment 5: Model Checking Paxos**  
  - Build a bare-bones model checker, apply it to Paxos, and analyze various consensus scenarios.  
  - Introduces the principles of formal modeling and model checking, which are increasingly important in distributed systems.  
