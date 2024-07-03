## Overview: Phone Number Normalizer
Phone numbers are read from a `.txt` file and added to an SQLite database. The phone numbers are then read from the database, normalized to the format `xxxxxxxxxx`, and written back to the database; any duplicates are removed.

This is a learning exercise. The goal is to better understand how interact with an SQLite database using Go. It is not meant to be an efficient or production-ready solution.

## Running the Program

1. Open a terminal at the root of the project. Build the project via the Makefile:
```bash
make
```

2. Run the program:
```bash
./phone_number_normalizer
```

