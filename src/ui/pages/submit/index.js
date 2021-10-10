/***************
 * Submit Route home page
 * 
 * Home page for the submit route.
 ***************/
import {React, Component} from 'react';
import Head from 'next/head'
import IconButton from '@mui/material/IconButton';
import Button from '@mui/material/Button';
import Icon from '@mui/material/Icon';
import Typography from '@mui/material/Typography';
import NavBar from '../../components/NavBar.js';
import Link from 'next/link';
import TextField from '@mui/material/TextField';
import MenuItem from '@mui/material/MenuItem';

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
        const formEntryStyle = {
            'width': '50%'
        }
        const fullLineStyle = {
            'width': '100%',
        }
        var products = [
            {
                value:'any',
                label:'Any'
            },
            {
                value:'devops-overflow',
                label:'DevOps Overflow'
            },
            {
                value:'stem-practice',
                label:'Stem Practice'
            }
        ]

        return (
            <div>
                <Head>
                    <title>Home Page</title>
                    <meta name="description" content="Submit Page" />
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
                    <div><Button variant="contained" color='primary'>Submit New Issue</Button></div>
                    <div style={formContainerStyle}>
                        <TextField  style={formEntryStyle} label={'Issue Name'} required = {true} varient='Outlined'></TextField>
                        <TextField  style={formEntryStyle} label={'Product'} select required={true} varient='Outlined'>
                            {products.map((option) => (
                                <MenuItem key={option.value} key={option.value}>
                                    {option.label}
                                </MenuItem>
                            ))}
                        </TextField>
                    </div>
                    <div style={formContainerStyle}>
                        <TextField style={fullLineStyle} label={'Issue Description'} required={true} multiline={true} minRows={2}  varient='Outlined'></TextField>
                    </div>
                </center>
            </div>
        )
    }
}
export default Home;
