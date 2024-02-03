# Introduction to datasebase storage

## 1. What is DBMS?

1. Stands for **D**ata**b**ase **M**anagement **S**ystem.
2. Responsible for managing data for your store (CRUD) operations on it.
3. Will try to optimize the queries if possible too.

### 1.1. DBMS Main components

![Alt](https://github.com/MohamedHassan499/today-i-learned/assets/50884619/1663f668-a97f-43a4-b4eb-8547711c9335)

## 2. Storage Manager

In disk-based DBMS, it tends to use non-volatile storages for storing huge among of data in a huge container rather than going for a volatile data storage for several reasons:

- Non-volatile storage means we keep the data even If unexpected circumstances happens, i.e database server crash, powering off. Unlike the volatile storage like **RAM**, which might resulting the stored data forever.
- Cheaper. We might want like 50TB sotrage container to store some massive data, doing this with volatile storages will be very very expensive, due to its **Random Access**.
- Non-volatile storages operates using **sequencial access**, meaning that It'd be better If we could access data in sequencial way, why? _because either HDD or SSD tend to move their head which points to the current disk block to read that specific memory block, and doing this is expensive and slower than accessing gthat storage unit with access memory_.
- Can access huge chunck of data at a time, unlike RAM at which we access bytes one by one.

## 2.1. How to get most out of our DBMS?

1. Try to save space as much as you can, you'll be storing huge amount of data.
2. Try to access data sequencially as much as you can, afterall, reading from the disk is expensive.
3. Try to minimize reads/write as much as you can.

## Components of reading/writing information on storage manager:

![enter image description here](https://github.com/MohamedHassan499/today-i-learned/assets/50884619/694e5fa9-ec22-4a55-ae1f-3c8a963d0a3f)

On executing query and reaching out to the storage manager:

1. **Query Engine** will know (Somehow) that the data we want to access exists in the page #3 in the disk. _Every page stored in the disk has its unique ID_.
2. It'll lookup in the memory buffer and see if that page exists or not.
3. If not, then we'll lookup in the disk, and we'll get the reference of the page containing the information we want.
4. In the end, the query result will be buffered in the memory and It'll return to the query engine the reference of that page in the disk.
   > 1. Memory buffer: Temporary storage in the RAM at which we **store most frequent data**
   > 2. The OS doesn't know the content of that page, **the interpreter** of the query engine does, and translates this to the data whatever the client wants

## 2.1 How data is stored

1. Every DBMS, has its own format for storing all the information, _and a specific format might not work for the other systems_
2. OS doesn't understand the content of the file, he sees everything as bits.
3. Every system might have one or more files depending on the system itself and when does it store.
4. These files contains pages **(Unit storage for the DBMS)**
5. ![enter image description here](https://github.com/MohamedHassan499/today-i-learned/assets/50884619/c4d121be-a2b3-40b3-88f1-025efdc5d859)

### 2.1.1 Pages:

1. Fixed size files _(8KB default for postgres, 16KB for MySQL)_
2. Files containing information like: - Metadata - Tuples (records of table) - indexes - Logs - Categories
3. Each page has its unique ID, at which DMBS can map a specific page ID to its physical address in the disk.

> 1. Storage manager keeps tracking the available space in the pages so it'll try to fit in some information there for better utilization for the disk space.
> 2. Data replication **is not** the responsibility of the storage manager, another component does that. It might be above or below the data storage component.

### 2.1.2. How pages are stored/organized?

![enter image description here](https://github.com/MohamedHassan499/today-i-learned/assets/50884619/ba250e1c-fe08-42e4-ad50-affe1f6fe581)

There are seveal ways for organizing files:

1. **Heap File Organization**: Used for storing tuples.
2. **Tree File Organization**: Used for indexing columns dealing with queries related to ranges `SELECT * FROM "User" WHERE ID >= 18 AND ID <= 24`
3. **Sequential/Sorted File Organization (ISAM)**: Used for querying on sequential data.
4. **Hashing File Organization**: For indexing columns with equality operating (Operates as hash table): `SELECT * FROM "User" WHERE ID = 18;`

### 2.1.2.1. Heap File Organization:

We agreed that on querying data based on some condition specified from the parser, the data will be stored inside pages, but we have 2 problems: 1. The pages might not be sorted. 2. Data inside the pages might now be sorted. 3. Pages might be distributed across many files depending on the DBMS.

What mechanism might work on looking for information in these files + pages?
Let's assume the best case scenario that these pages are sorted in one file. We'd do this (Knowing the size of the pages:)
<center>$( \text{Offset} = \text{\#Page} \times \text{Page Size})$</center>

But what If we had multiple files? This algorithm won't work.
![enter image description here](https://github.com/MohamedHassan499/today-i-learned/assets/50884619/05744273-9cee-4602-a3c6-6cf4234cf34f)

I'll need to store extra meta-data so it'd help me for the lookup of the page I want to access. Here comes the:

### Heap Files - Page directory :

1. Contains a special page(s) at the begenning of the file containing:

- Location for each page in the file
- Number of free slots per page
- List free/empty pages

This page must be in sync with the data pages

![enter image description here](https://github.com/MohamedHassan499/today-i-learned/assets/50884619/fd40cb7b-5d95-419e-8b55-20dca37a85ae)
