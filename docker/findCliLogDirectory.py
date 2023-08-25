# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

import os
import sys

run_prefix="network_"
restart_prefix="restart_"


def get_datetime(prefix, parent_dir):
	max_date=0
	max_time=0
	max_date_str=""
	max_time_str=""
	for x in os.listdir(os.path.expanduser('~') +"/.avalanche-cli/runs" + parent_dir):
		if not x.startswith(prefix):
			continue
		trimmed=x[len(prefix):]
		date_arr=trimmed.split("_")
		date=int(date_arr[0])
		time=int(date_arr[1])
		if date > max_date:
			max_date = date
			max_time = time
			max_date_str = date_arr[0]
			max_time_str = date_arr[1]
		elif date == max_date and time > max_time:
			max_time = time
			max_time_str = date_arr[1]

	return prefix+max_date_str+"_"+max_time_str

def main(argv):
	path_str=""
	run_parent_dir=""
	if len(argv) > 0 and argv[0] == "restart":
		restart_datetime=get_datetime(restart_prefix, "")
		path_str+=restart_datetime+"/"
		run_parent_dir="/"+path_str
	run_datetime=get_datetime(run_prefix,run_parent_dir)
	path_str+=run_datetime
	print(path_str)

if __name__ == "__main__":
	main(sys.argv[1:])