#include "ray.h"
#include "raylib.h"
#include <stdlib.h>

int posX = 0;
int posY = 0;

int screenWidth = 800;
int screenHeight = 450;

void setPos(int x, int y) {
    posX = x;
    posY = y;
}

int renderLoop(void)
{
    InitWindow(screenWidth, screenHeight, "Pong");
    SetTargetFPS(60);

    while (!WindowShouldClose())
    {
        GoLogicLoop();
        BeginDrawing();
            ClearBackground(RAYWHITE);
            GoRenderLoop();
            DrawFPS(screenWidth - 100, 20);
        EndDrawing();
    }
    CloseWindow();

    return 0;
}
