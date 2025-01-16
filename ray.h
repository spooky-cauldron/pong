#pragma once

// #include "_cgo_export.h"

typedef void (*UpdateFunc) ();
int renderLoop(void);
void logicLoop(void);
void addUpdateFunc(UpdateFunc f);
void setPos(int x, int y);
extern void GoLogicLoop();
extern void GoRenderLoop();
// int ScreenWidth();
// int ScreenHeight();
