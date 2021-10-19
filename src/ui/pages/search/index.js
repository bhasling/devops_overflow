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
import RestCall from '../../components/RestCall.js'
import Typography from '@mui/material/Typography';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';

class Home extends Component {
    constructor(props) {
         super();
         this.props = props;
         this.state = {
             isLoaded: false,
             issues: [],
             errorMessage: "Enter search keyword(s) and click Search."
         }
    }

    clickSearch() {
        var url = `./issues`;
        RestCall.invoke("GET", url , null, "Unable to load issues.", this.simulateSearch())
        .then (
            (response) => this.getIssuesResponse(response),
            (message) => this.errorResponse(message));
    }

    getIssuesResponse(response) {
        var message = `Found ${response.length} issues.`
        this.setState({isLoaded:true, issues:response, errorMessage: message});
    }

    clearMessage() {
        this.setState({errorMessage: ""});
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

                    <div><Button onClick={()=>this.clickSearch()} variant="contained" color='primary'>Search</Button></div>
                    <div onClick={()=>this.clearMessage()}>
                        <Typography fontWeight = 'bold' fontSize='1.5rem' color='secondary'>
                            {this.state.errorMessage}
                        </Typography>
                    </div>

                    <div style={formContainerStyle}>
                        <TextField style={fullLineStyle} label={'Search keywords'} required = {true} varient='Outlined'></TextField>
                    </div>
                </center>
                <TableContainer component={Paper} sx={{marginRight:"20px", marginLeft:"20px"}}>
                    <Table sx={{marginLeft:"20px", marginRight:"20px"}} aria-label='List of Issues'>
                        <TableHead>
                            <TableRow>
                                <TableCell>Issue Title</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                        {
                            this.state.issues.map((issue) => (
                                <TableRow key={issue.issue_id} sx={{cursor:"pointer"}}>
                                    <TableCell>{issue.title}</TableCell>
                                </TableRow>
                            ))
                        }
                        </TableBody>
                    </Table>
                </TableContainer>
            </div>
        )

    }

    simulateSearch() {
        var result = null;
        if (process.env.NODE_ENV == "development") {
            var jsonString = `
            [
                {
                    "issue_id": "126ad688-9b9d-4680-adf1-f99831b8b677",
                    "user_id": "",
                    "title": "My first issue",
                    "product": "",
                    "SubmitTime": "0001-01-01T00:00:00Z",
                    "desciption": "",
                    "answers": [
                        {
                            "AnswerId": "ad4b7842-7bec-4695-a0e9-247222ef74b4",
                            "user_id": "",
                            "AnswerTime": "0001-01-01T00:00:00Z",
                            "desciption": ""
                        }
                    ]
                },
                {
                    "issue_id": "126ad688-9b9d-4680-adf1-f99831b8b678",
                    "user_id": "",
                    "title": "My Second issue",
                    "product": "",
                    "SubmitTime": "0001-01-01T00:00:00Z",
                    "desciption": "",
                    "answers": [
                        {
                            "AnswerId": "ad4b7842-7bec-4695-a0e9-247222ef74b4",
                            "user_id": "",
                            "AnswerTime": "0001-01-01T00:00:00Z",
                            "desciption": ""
                        }
                    ]
                }
            ]
            `;
            result = JSON.parse(jsonString);
        }
        return result;
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

