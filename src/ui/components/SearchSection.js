/***************
 * Search Route home page.
 * This page allows a user to search for issues by keywords and view the list of found issues.
 * The user can click on an issue to view the details and the answers of the issue.
 * 
 ***************/
import {React, Component} from 'react';
import { connect } from "react-redux";
import IconButton from '@mui/material/IconButton';
import Icon from '@mui/material/Icon';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Link from 'next/link';
import RestCall from './RestCall.js'
import Typography from '@mui/material/Typography';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';

class SearchSection extends Component {
    constructor(props) {
         super();
         this.props = props;
         this.state = {
             isLoaded: false,
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
        this.props.updateSearchResult(response);
        //this.setState({isLoaded:true, errorMessage: message});
    }

    errorResponse(message) {
        this.setState({errorMessage: message})
    }

    clearMessage() {
        this.setState({errorMessage: ""});
    }
    clickIssue(issue) {
        this.props.setIssue(issue);
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
                <TableContainer component={Paper} style={{marginLeft:"20px", width:"95%"}}>
                    <Table aria-label='List of Issues'>
                        <TableHead>
                            <TableRow>
                                <TableCell sx={{fontWeight:"bold"}}>Issue Title</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                        {
                            this.props.searchResult.map((issue) => (
                                <TableRow key={issue.issue_id} sx={{cursor:"pointer"}}>
                                    <TableCell onClick={()=>this.clickIssue(issue)}>{issue.title}</TableCell>
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
                    "product": "DevOps",
                    "SubmitTime": "0001-01-01T00:00:00Z",
                    "description": "this is description",
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
    updateSearchResult: (searchResult) => dispatch({type: 'updateSearchResult', searchResult:searchResult})
   }
}
export default connect(mapStateToProps, mapDispatchToProps)(SearchSection);

