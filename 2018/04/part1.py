from datetime import datetime

events = {}
with open('input.txt', 'r') as infile:
	for line in infile.readlines():
		time = line.split(']')[0][1:]
		message = line.split(']')[1][1:-1]

		timestamp = datetime.strptime(time, '%Y-%m-%d %H:%M')
		events[timestamp] = message

types = {'new': 'Guard', 'wake': 'wakes', 'sleep': 'falls'}

guard_id = 0
minutes_asleep = 0
sleep_start = None

guards_total_sleep = {}
guards_sleep_by_min = {}

for time in sorted(events.keys()):
	message = events[time].split(' ')
	msg_type = message[0]

	if msg_type == types['new']:
		guard_id = int(message[1][1:])

		if guard_id not in guards_total_sleep:
			guards_total_sleep[guard_id] = minutes_asleep

		if guard_id not in guards_sleep_by_min:
			guards_sleep_by_min[guard_id] = {}

			for i in range(60):
				guards_sleep_by_min[guard_id][i] = 0

		minutes_asleep = 0

	elif msg_type == types['sleep']:
		sleep_start = time
	elif msg_type == types['wake']:
		minutes_asleep += (time - sleep_start).total_seconds() // 60
	else:
		print('Invalid message')

most_asleep = max(guards_total_sleep, key=guards_total_sleep.get)
print('#%d: %d minutes asleep' % (most_asleep, guards_total_sleep[most_asleep]))
