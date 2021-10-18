/***************
 * NavBar.js
 * 
 * Implements a <NavBar> tag that display a material design nav bar.
 */
import React from 'react'
import AppBar from '@mui/material/AppBar'
import Toolbar from '@mui/material/Toolbar'
import Typography from '@mui/material/Typography'
import SignInButton from './SignInButton.js';
import IconButton from '@mui/material/IconButton';
import Icon from '@mui/material/Icon';
import Link from 'next/link';

const NavBar = () => {
    var iconStyle = {
        'fontSize':'20px',
        'display': 'inline-block',
        'width': '20px',
        'marginTop' : '4px',
        'verticalAlign':'top',
        'fontFamily': 'Material Icons',
        'cursor':'pointer'
    }
    return(
        <AppBar position='static'>
            <Toolbar >
                <Typography color="inherit">
                    Dev Ops Overflow
                </Typography>
                <div style={{'marginLeft': "auto"}}>
                    <Link href='/configure' passHref>
                    <IconButton style={{'verticalAlign': 'top', 'marginRight': '10px', 'color': 'white'}} >
                        <Icon >settings</Icon>
                    </IconButton>
                    </Link>
                    <SignInButton />
                </div>
            </Toolbar>
        </AppBar>
    )
}
export default NavBar;

