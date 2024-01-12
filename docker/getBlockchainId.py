# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

import sys
import json

def get_id_by_name(data, name):
    for obj in data:
        if obj['name'] == name:
            return obj['id']
    return None

def main(argv):
    data = json.load(sys.stdin)
    name = argv[0]
    print(get_id_by_name(data, name))

if __name__ == "__main__":
	main(sys.argv[1:])