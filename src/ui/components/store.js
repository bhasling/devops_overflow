/*********************
 * store.js
 *
 * Implements the Redox store, created by createStore with a reducer function.
 * The reducer function reduces an action to update the redux state.
 */
 import {createStore} from 'redux';

 function storeReducer(state = [], action) {
    switch (action.type) {
        case 'updateLoginState':
            return {
                ...state,
                loginState : action.loginState
            }
        default:
            return state;
     }
 }
 
 var store = createStore(storeReducer, {
     "version": 4,
     'loginState': {
         'loginName': null,
         'loginId': null
     }
 });
 
 export default store
 
 