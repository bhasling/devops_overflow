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
import RestCall from '../../components/RestCall.js';
import Typography from '@mui/material/Typography';
import NavBar from '../../components/NavBar.js';
import Link from 'next/link';
import TextField from '@mui/material/TextField';
import MenuItem from '@mui/material/MenuItem';

class SubmitPage extends Component {
    constructor(props) {
         super();
         this.props = props;
         this.state = {
             product: "",
             title: "",
             description: "",
             errorMessage: ""

         }
    }

    changeProduct(event) {
        this.setState({product: event.target.value})
    }
    changeTitle(event) {
        this.setState({title: event.target.value})
    }
    changeDescription(event) {
        this.setState({description: event.target.value})
    }
    clearMessage() {
        this.setState({errorMessage: ""});
    }

    clickSubmit() {
        var message = "";
        if (this.state.title == "") {
            message = message + "Title is required. "
        }
        if (this.state.product == "") {
            message = message + "Product is required. "
        }
        if (this.state.description == "") {
            message = message + "Description is required. "
        }
        this.setState({errorMessage: message});
        if (message == "") {
            // No errors so submit the new issue
            var url = `./api/issues`;
            var body = {
                title: this.state.title,
                product: this.state.product,
                description: this.state.description
            }
            RestCall.invoke("POST", url , body, "Unable to submit issue.", this.simulateSubmit())
            .then (
                (response) => this.getIssuesResponse(response),
                (message) => this.errorResponse(message));
            }
    }

    getIssuesResponse(response) {
        var message = `Submitted new issue.`
        this.setState({errorMessage: message});
    }

    errorResponse(message) {
        this.setState({errorMessage: message})
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
                    <Link href="/" passHref>
                        <IconButton style={{'verticalAlign': 'top', 'marginRight': '10px'}} color='primary'>
                            <Icon style={{'fontSize':'2rem'}}>home</Icon>
                        </IconButton>
                    </Link>
                    <div>
                        <Button onClick={()=>this.clickSubmit()} variant="contained" color='primary'>
                                    Submit New Issue
                        </Button>
                    </div>
                    <div onClick={()=>this.clearMessage()}>
                        <Typography fontWeight = 'bold' fontSize='1.5rem' color='secondary'>
                            {this.state.errorMessage}
                        </Typography>
                    </div>
                    <div style={formContainerStyle}>
                        <TextField  style={formEntryStyle} 
                                    label={'Issue Title'}
                                    onChange={(event)=>this.changeTitle(event)}
                                    required = {true} 
                                    varient='Outlined'></TextField>
                        <TextField  style={formEntryStyle}
                                    label={'Product'}
                                    select 
                                    required={true} 
                                    varient='Outlined'
                                    onChange={(event)=>this.changeProduct(event)}
                                    value={this.state.product}>
                            {products.map((product) => (
                                <MenuItem key={product.value} value={product.label}>
                                    {product.label}
                                </MenuItem>
                            ))}
                        </TextField>
                    </div>
                    <div style={formContainerStyle}>
                        <TextField style={fullLineStyle} 
                                   label={'Issue Description'} 
                                   required={true} 
                                   onChange={(event)=>this.changeDescription(event)}
                                   multiline={true} 
                                   minRows={2}  
                                   varient='Outlined'></TextField>
                    </div>
                </center>
            </div>
        )
    }

    simulateSubmit() {
        var result = null;
        if (process.env.NODE_ENV == "development") {
            var jsonString = `
            {
                "message": "Submitted"
            }
            `;
            result = JSON.parse(jsonString);
        }
        return result;
    }

}
export default SubmitPage;
