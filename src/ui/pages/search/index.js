/***************
 * Search Route home page
 * 
 * Home page for the submit route.
 ***************/
import {React, Component} from 'react';
import Head from 'next/head'
import { connect } from "react-redux";
import IconButton from '@mui/material/IconButton';
import Icon from '@mui/material/Icon';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import NavBar from '../../components/NavBar.js';
import Link from 'next/link';

class Home extends Component {
    constructor(props) {
         super();
         this.props = props;
    }

    render() {
        const formContainerStyle = {
            'display': 'flex',
            'columnGap': '30px',
            'margin': '20px'
        }

        const fullLineStyle = {
            'width': '100%',
            'margin':  '20px'
        }

        return (
            <div>
                <Head>
                    <title>Home Page</title>
                    <meta name="description" content="Searh Page" />
                    <link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons"/>
                    <link rel="icon" href="/favicon.ico" />
                </Head>
                <NavBar/>
                <center>
                    <Link href="/" passHref>
                        <IconButton style={{'verticalAlign': 'top', 'marginRight': '10px'}} color='primary'>
                            <Icon style={{'fontSize':'2rem'}}>home</Icon>
                        </IconButton>
                    </Link>

                    <div><Button variant="contained" color='primary'>Search</Button></div>
                    <div style={formContainerStyle}>
                        <TextField style={fullLineStyle} label={'Search keywords'} required = {true} varient='Outlined'></TextField>
                    </div>

                </center>
            </div>
        )
    }
}
const mapStateToProps = state => {
    var props = {loginState: state.loginState, searchResult : state.searchResult};
    return props;
}
const mapDispatchToProps = (dispatch) => {
   return {
    updateSearchResult: (searchResult) => dispatch({type: 'updateSearchResult', searchResult})
   }
}
export default connect(mapStateToProps, mapDispatchToProps)(Home);

