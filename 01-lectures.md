---
layout: page
title: Lectures
---

The following topics will be presented over the course of the semester.
Each topic will be covered in (roughly) one lecture.
Lecture notes are linked as they become available.

  0. [Course introduction]({{ site.baseurl
     }}/lectures/00-course-introduction.pdf)
      * CAs' homework series intro (on courseworks/Files)
      * Go tutorial (on courseworks/Files)

  1. [Distributed systems primer]({{ site.baseurl }}/lectures/01-distributed-systems-primer.pdf)
      * challenges and goals of distributed systems
      * example architectures
      
  2. [Distributed computation]({{ site.baseurl }}/lectures/02-map-reduce.pdf)
      * MapReduce
      * Spark
      * Tradeoffs

  3. [Communication models]({{ site.baseurl }}/lectures/03-rpc.pdf)
      * remote procedure calls (RPC)
      * RPC libraries
      * failure models
      * semantics

  4. [Time and coordination]({{ site.baseurl }}/lectures/04-clocks.pdf)
      * challenges
      * physical and logical clocks
      * distributed mutual exclusion

  5. [Agreement in distributed systems]({{ site.baseurl
     }}/lectures/05-agreement-problem.pdf)
      * the atomic commitment problem
      * the consensus problem
      * use cases for each
      * FLP impossibility result of achieving consensus

  6. [The transaction abstraction]({{ site.baseurl
     }}/lectures/06-local-transactions.pdf)
      * ACID semantics
      * concurrency control mechanisms
      * recovery mechanisms

  7. [Atomic commitment protocols]({{ site.baseurl }}/lectures/07-2pc.pdf)
      * 2-phase-commit
      * blocking nature

  8. [Consensus protocols]({{ site.baseurl }}/lectures/08-paxos.pdf)
      * Paxos overview, key ideas, basic algorithm
      * examples of normal operation and operation under failures
      * liveness failure mode
      * multi-Paxos
      * applications

  9. Case studies from industry:
      * [Google's Spanner scalable, fault-tolerant ACID database]({{ site.baseurl }}/lectures/09-spanner.pdf)
      * [Google's Chubby fault-tolerant lock service]({{ site.baseurl
      }}/lectures/09-chubby.pdf)
      * [Google's Bigtable scalable, fault-tolerant, multi-dimensional, sorted
      map]({{ site.baseurl }}/lectures/09-bigtable.pdf)

  10. [Broader view of isolation and consistency
     semantics]({{ site.baseurl }}/lectures/10-broader-semantics.pdf)
      * isolation: serializability, repeatable reads, read committed, read uncommitted
      * consistency: external, sequential, causal, eventual
      * mechanisms for each
      * performance/usability tradeoffs

  11. [Beyond storage and MapReduce: Broader infrastructure systems]({{ site.baseurl }}/lectures/11-large-scale-software-stacks.pdf)
      * Google's software stack
      * Meta's software stack
      * Hadoop and Spark software stacks

  12. [Cluster scheduling]({{ site.baseurl }}/lectures/12-scheduling.pdf)
      * scheduler architectures and considerations
      * frameworks: YARN, Mesos, Borg
      * algorithms: dominant resource fairness, bin packing

  13. [Testing and model checking]({{ site.baseurl
      }}/lectures/13-testing-model-checking.pdf)
      * testing approaches and challenges
      * formal specification and model checking
      * TLA+ primer

  14. [Security and Byzantine fault tolerance]({{ site.baseurl }}/lectures/14-bft.pdf)

