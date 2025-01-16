#include "ray.h"
#include "raylib.h"
#include <stdlib.h>


// UpdateFunc* updateFuncs;
// int nUpdateFuncs = 0;

int posX = 0;
int posY = 0;

int screenWidth = 800;
int screenHeight = 450;

// int ScreenWidth() {
//     return screenWidth;
// }

// int ScreenHeight() {
//     return screenHeight;
// }

void setPos(int x, int y) {
    posX = x;
    posY = y;
}

int renderLoop(void)
{
    // int updateFuncCount = 1024;
    // updateFuncs = (UpdateFunc*) malloc(updateFuncCount * sizeof(UpdateFunc));

    InitWindow(screenWidth, screenHeight, "Pong");
    SetTargetFPS(60);

    while (!WindowShouldClose())
    {
        // logicLoop();
        GoLogicLoop();
        BeginDrawing();
            ClearBackground(RAYWHITE);
            GoRenderLoop();
            // DrawText("Congrats! You created your first window!", 190, 200, 20, LIGHTGRAY);
            // DrawCircle(posX, posY, 20.0f, GREEN);
            DrawFPS(screenWidth - 100, 20);
        EndDrawing();
    }

    CloseWindow();

    return 0;
}



// void addUpdateFunc(UpdateFunc f) {
//     updateFuncs[nUpdateFuncs] = f;
//     nUpdateFuncs++;
// }

// void logicLoop(void) {
//     if (nUpdateFuncs == 0) {
//         return; 
//     }

//     for (int i = 0; i < nUpdateFuncs; i++) {
//         updateFuncs[i]();
//     }
// }