#include "command_handler.h"
#include <stdio.h>
#include <stdlib.h>

void create_dir(const char *name) {
    char command[256];
    sprintf(command, "mkdir -p %s", name);
    system(command);
}

void create_file(const char *filename, const char *content) {
    FILE *file = fopen(filename, "w");
    if (file) {
        fprintf(file, "%s", content);
        fclose(file);
        printf("File created %s\n", filename);
    } else {
        printf("Failed to create file: %s\n", filename);
    }
}

void create_project(const char *name) {
    printf("Creating a new GoRails project called %s...\n", name);
    char command[256];
    sprintf(command, "mkdir -p %s/{controllers,models,views,routes}",name);
    system(command);
    create_dir("controllers");
    create_dir("views");
}

void run_server() {
    printf("Starting the server...\n");
    system("go run index.go");
}

void generate_demo(const char *name) {
    printf("Generating demo for %s...\n", name);
    char dir_command[256];
    sprintf(dir_command, "%s", name);
    create_dir(dir_command);
    create_dir("controllers");

    char file_path[256];
    sprintf(file_path, "%s/index.go", name);

    create_file(file_path, "package main\n\n"
                         "import (\n"
                         "\t\"fmt\"\n"
                         "\t\"gorails/core/router\"\n"
                         "\t\"net/http\"\n"
                         ")\n\n"
                         "func main() {\n"
                         "\tr := router.NewRouter()\n\n"
                         "\t// Static route\n"
                         "\tr.GET(\"/\", func(w http.ResponseWriter, r *http.Request, _ map[string]string) {\n"
                         "\t\tfmt.Fprintln(w, \"<h1>Welcome to GoRails!</h1>\")\n"
                         "\t})\n\n"
                         "\tfmt.Println(\"Server is running on port 8000\")\n"
                         "\thttp.ListenAndServe(\":8000\", r)\n"
                         "}\n");

    
    printf("GoRails: Demo %s created successfully!\n", name);
    printf("Run `gorails start` to start the server\n");
}

