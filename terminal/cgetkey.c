#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/time.h>
#include <termios.h>

static struct termios g_old_kbd_mode;

static int kbhit(void) {
	struct timeval timeout;
	fd_set read_handles;
	int status;

	FD_ZERO(&read_handles);
	FD_SET(0, &read_handles);
	timeout.tv_sec = timeout.tv_usec = 0;
	status = select(0 + 1, &read_handles, NULL, NULL, &timeout);
	return status;
}

int readkey( void ) {
	int ch;
	struct termios new_kbd_mode;

	tcgetattr(0, &g_old_kbd_mode);
	memcpy(&new_kbd_mode, &g_old_kbd_mode, sizeof(struct termios));

	new_kbd_mode.c_lflag &= ~(ICANON | ECHO);
	new_kbd_mode.c_cc[VTIME] = 0;
	new_kbd_mode.c_cc[VMIN]  = 1;

	tcsetattr(0, TCSANOW, &new_kbd_mode);

	while (!kbhit())
	{
		ch = getchar();
		tcsetattr(0, TCSANOW, &g_old_kbd_mode);
		return ch;
	}

	tcsetattr(0, TCSANOW, &g_old_kbd_mode);
	return ch;
}