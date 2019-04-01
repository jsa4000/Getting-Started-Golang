# Data Generator

## Values

- RowsPerPartition = 100000 rows (size?) // Maximum rows per partition (process), limiting the memory and overhead
- RowsPerCommit = 100 // Maximun rows per commit in bulk operation. Bigger values, worst results if the process fails.
- ChunksPerPartition = 100000 / chunks? = RowsPerCommit (chunk)  // This is the number of chunks or iterations per partition (max)

## Computation (split process)

RowsTotal = 10000000 (3.25GB)
Partitions = 10000000 / RowsPerPartition (100000) = 100 partitions

## Split process

The split process will consist, first in generate the file and them split the file into chunks.

The file is split by lines, since it is the minimal logic unit, and it cannot be split into different files to be processed.

```bash

docker run -it -v /tmp/test:/tmp/test busybox /bin/sh

cd /tmp/test

split -l 100000 -d -a 3 file.csv split_file_

```

| Time | Start | End | Processes |
| -- | -- | -- | -- |
| | 20:32 | 20:45 | (1.2GB)  |
| | 20:32 | 21:07 | (3.25GB) |

Parallel split??

The process must generate the `md5sum` in order to verify the integrity of the files split.
https://askubuntu.com/questions/318530/generate-md5-checksum-for-all-files-in-a-directory