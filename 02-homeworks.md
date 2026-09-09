---
layout: page
title: Homeworks/Exams
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
**48-hour grace period** per assignment with a total of **72-hour grace period** accumulated over all homeworks, for which you will not be
downgraded.  Lateness is accounted at hour granularity (i.e., 1 second late == 1 hour late).
Once you reach 48 hour on a single assignment or 72 hours total of lateness, the due homework to incur even a one-second delay
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


## Grading of Homeworks

Please refer to the homeworks and deadlines for the assignments.
The Homeworks are 8 individually graded assignments ( note hw2,3,4 have parts a and b ). We take the best 6 of 7 of homeworks 1 - 4b for 10% each ( = 60%). Homework 5 is mandatory with also 10%. Hence your homeworks will contribute 70% of your grade.

## Code Walkthrough

To ensure the authenticity of your code we will have multiple 1:1 code walkthrough (you walk we ask questions). These are meant to ascertain you actually wrote and understand your code. These will be short and nothing to get excited about if you handed in original code. We will allow later in class the frequency and timing.

## Final Exam

There is an in person final exam on Mon 12/21 at 7:00pm (exception to the "in person" ofc is for CVN students).
The exam counts 30% of your final grade.
