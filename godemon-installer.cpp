#include <iostream>
#include <string>

const std::string version = "21.04-beta";

void buildApp(){
    std::string temp ="unzip " + version + ".zip"; 
    system(temp.c_str());
    temp = "cd ./godemon-" + version;
    system(temp.c_str());
    system("go build");
}

void prepareDirs(){
    system("mkdir ~/.godemon");
    system("mkdir ~/.godemon/logs/");
    system("mkdir ~/.godemon/bin/");
    system("git clone https://github.com/Godemon-simplify-your-Go-programming/Godemon-update");
    system("cd Godemon-update");// to do 
    system("g++ godemon_update.cpp -o godemon-update");
    system("cp godemon-update ../");
    system("sudo chmod 777 ./godemon-update");
    system("mv ./godemon-update ~/.godemon/bin");
}

void removing(){
    system("cd ..");
    std::string temp = "sudo rm -r godemon-"+ version;
    system(temp.c_str());
    temp = "sudo rm -r "+version+".zip";
    system(temp.c_str()); 
}

int main(){
    int option=0;
    int GL=0;
    std::cout<<"*************__*************\n************/*/*************\n***********/*/**************\n**********/*/____***********\n*********/____**/***********\n*************/*/************\n************/*/*************\n***********/*/**************\n**********/*/***************\n*********/_/****************\n****************************\n";
    std::cout<<"\nWhat do you want to do? \n1. Install Godemon \n2. Update Godemon \n";
    std::cout<<"\nAnswer: ";
    std::cin>>option;
    std::cout<<"\nDo you want to do this: \n1. Global \n2. Local \n";
    std::cout<<"\nAnswer: ";
    std::cin>>GL;

    std::string wget = "wget https://github.com/nProgrammer/godemon/archive/"+version+".zip";

    system(wget.c_str());
    if(option == 1 ){
        if(GL==2){
            buildApp();
            prepareDirs();
            system("sudo mv godemon ~/.godemon/bin/");
            system("sudo chmod 777 ~/.godemon/bin/godemon");
            removing();
        }
        else if(GL==1){
            buildApp();
            system("sudo mkdir /usr/local/.godemon");
            prepareDirs();
            system("sudo mkdir /usr/local/.godemon/bin/");
            system("sudo mv godemon /usr/local/.godemon/bin/");
            system("sudo chmod 777 /usr/local/.godemon/bin/godemon");
            removing();
        }
    }
    else if(option == 2){
        if(GL==2){
            buildApp();
            system("sudo mv godemon ~/.godemon/bin/");
            system("sudo chmod 777 ~/.godemon/bin/godemon");
            removing();
        }
        else if (GL==1){
            buildApp();
            system("sudo mv godemon /usr/local/.godemon/bin/");
            system("sudo chmod 777 /usr/local/.godemon/bin/godemon");
            removing();
        }
    }

    std::cout<<"Everything done"<<std::endl;
    if(option == 1){
        if(GL == 2) {
            std::cout<<"\nNow add to .bashrc or .zshenv following line: export PATH=\$PATH:~/.godemon/bin \n";
        }
        else if(GL == 1){
            std::cout<<"\nNow add to .bashrc following line: export PATH=\$PATH:/usr/local/.godemon/bin \nIf it's MacOS - add this line to .zshenv and export PATH=\$PATH:~/.godemon/bin \n";
        }
    }
} 