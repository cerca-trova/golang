#include <stdio.h>
#include <string.h>

struct Books
{
    char title[50];
    char author[100];
    char subject[100];
    int  id;
};

void printBook(struct Books *book);
int main(){
    struct Books Book1;
    /* struct Books Book2; */

    strcpy(Book1.title, "C in chen");
    Book1.id = 123456;
    strcpy(Book1.subject,"hello world");
    strcpy(Book1.author,"chenran");

    printBook(&Book1);

    return 0;
}
void printBook(struct Books *book){
    printf("author is: %s\n",book->author);
    printf("id is: %d\n",book->id);
    printf("subject is: %s\n",book->subject);
    printf("title is: %s\n",book->title);
}