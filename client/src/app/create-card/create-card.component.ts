import { Component } from '@angular/core';
import axios from 'axios';
import { MatDialog, MatDialogRef } from '@angular/material/dialog';
import { HttpClient } from '@angular/common/http';

import {Router} from '@angular/router';



@Component({ templateUrl: 'create-card.component.html',
styleUrls: ["./create-card.component.css"] })

export class CreateCardComponent {

    constructor(private router: Router) {}

    frontItems: string[] = [''];
    backItems: string[] = [''];
  
    createCards() {
      

      const headers = { Authorization: sessionStorage.getItem('accessToken') };
      
      for (var i = 0; i < this.frontItems.length; i++) {

        const data = { 
            folderId: localStorage.getItem("folderId"),
            question: this.frontItems[i],
            answer: this.frontItems[i]
         };

      axios.post('http://api.memorly.kro.kr/users/card', data, { headers })
          .then(response => {
            // Request was successful, log the response data
            console.log(response.data);
            if (i == this.frontItems.length - 1) {
                this.router.navigateByUrl('card-view');
            }

          })
          .catch(error => {
            // Request failed, log the error message
            console.error(error.message);
          });
        }
    }
  
    addItem() {
      this.frontItems.push('');
      this.backItems.push('');
    }
  
    removeItem() {
      this.frontItems.pop();
      this.backItems.pop();
    }


    /*
    //cards:  = [];
    items: string[] = ['','','','','',''];

    addItem() {
        this.items.push('');
      }
    
      removeItem() {
        this.items.pop();
      }

    createCards(folderName1 : string,  cards1 : Array<Array<string>>) {

        const data = { 
          folderName: folderName1, 
          cards: cards1
          //for (let index = 0; index < cards1.length; index++) {
          //  cards[index][0] = cards1[index][0];
          //  cards[index][1] = cards1[index][1];
          //}  
          };

        // Define the headers with the access token
        const headers = { Authorization: sessionStorage.getItem('accessToken') };
        
        // Make the POST request
        axios.post('http://api.memorly.kro.kr/users/create-card', data, { headers })
          .then(response => {
            // Request was successful, log the response data
            console.log(response.data);

            this.router.navigateByUrl('library');

          })
          .catch(error => {
            // Request failed, log the error message
            console.error(error.message);
          });



    }
*/
}