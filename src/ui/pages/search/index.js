/***************
 * Search Route home page
 * 
 * Home page for the submit route.
 ***************/
import {React, Component} from 'react';
import Head from 'next/head'
import IconButton from '@mui/material/IconButton';
import Icon from '@mui/material/Icon';
import Typography from '@mui/material/Typography';
import NavBar from '../../components/NavBar.js';
import Link from 'next/link';

class Home extends Component {
    constructor(props) {
         super();
         this.props = props;
    }

    render() {
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
                    <Link href="/">
                        <IconButton style={{'verticalAlign': 'top', 'marginRight': '10px'}} color='primary'>
                            <Icon style={{'fontSize':'2rem'}}>home</Icon>
                        </IconButton>
                    </Link>
                    <Typography color='primary' fontWeight = 'bold' fontSize = '2rem'>
                    Search Page
                    </Typography>
                </center>
            </div>
        )
    }
}
export default Home;
