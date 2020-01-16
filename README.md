# cargo

This is for persistent storage.  
Requires a mechanism to for global information access.  
Persistence means either:
- Saving the data changes as a container
- Keeping the data always alive somewhere  

The first requires some sort of container hub, which were it to be internal, means that it is essentially doing the latter, but in a simplified way.

Investigate the alternatives: Kubeedge, swarmkit, other moby tech.

Anticipated: A docker-only version of the system is inferior to these others, but a wider version specifically designed for consumer results in a specialization for the system.

See how these others manage data, and build set-up a simple system for data management. 
