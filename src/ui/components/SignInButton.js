/****************
 * SignInButton.js
 * 
 * Displays the name of the logged in user or a SignIn button to login.
 * The SignIn button shows a login dialog where the user can login. This
 * post a request to the server with a user/password and returns with
 * the name of the user and an authentication cookie to allow other REST
 * requests to work.
 * 
 */
import {React, Component} from 'react';
import { connect } from "react-redux";
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import Typography from '@mui/material/Typography';
import RestCall from './RestCall.js'

class SignInButton extends Component {
     constructor(props) {
         super();
         this.props = props;
         this.state = {
             loginOpen: false,
             signUpOpen: false,
             signupUser: "",
             signupPassword: "",
             loginUser: "",
             loginPassword: "",
             message: ""
         }
     }
 
 
     render() {
        var iconStyle = {
            'fontSize':'20px',
            'display': 'inline-block',
            'width': '20px',
            'marginTop' : '4px',
            'verticalAlign':'top',
            'fontFamily': 'Material Icons',
            'cursor':'pointer'
        }
        var nameStyle = {
            'fontSize':'16px',
            'marginTop' : '4px',
            'verticalAlign':'middle',
            'display': 'inline-block',
            'fontWeight': 'bold',
            'cursor':'pointer'
        }
        var signUpoutStyle = {
            'cursor':'pointer'
        }
        if (this.props.loginState.userHover) {
            nameStyle['textDecoration'] = 'underline';
        }
        if (this.props.loginState.signUpOutHover) {
            signUpoutStyle['textDecoration'] = 'underline';
        }
        return(
            <div style={{'display':'inline-block', 'textAlign': 'left'}}                
            >
                <div>
                    <span>
                        <span style={iconStyle}>{this.props.loginState.userHover ? 'person_outline' : 'person'}</span>
                        <span style={{'display':'inline-block'}}>
                            <div style={nameStyle} onClick={()=>this.clickPeople()} onMouseEnter={()=>this.userMouseEnter()} onMouseLeave={()=>this.userMouseLeave()}>
                                {this.props.loginState.loginId? this.props.loginState.loginUser: 'Sign In'}
                            </div>
                            <div style={signUpoutStyle}  onClick={()=>this.clickSignUpOut()} onMouseEnter={()=>this.signOutUpMouseEnter()} onMouseLeave={()=>this.signOutUpMouseLeave()}>
                                {this.props.loginState.loginId? 'Sign Out': 'Sign Up'}
                            </div>
                        </span>
                    </span>

                    <Dialog open={this.state.loginOpen} onClose={()=>this.closeLoginDialog()}>
                        <DialogTitle>Login</DialogTitle>
                        <DialogContent>
                            <DialogContentText>
                                <Typography fontSize='1rem'>
                                    {this.state.message}
                                </Typography>                               
                            </DialogContentText>
                            <TextField
                                autoFocus
                                margin="dense"
                                id="name"
                                label="User Id"
                                type="string"
                                onChange={(event)=>this.changeLoginUser(event)}
                                fullWidth
                                variant="standard"
                            />
                            <TextField
                                margin="dense"
                                id="password"
                                label="Password"
                                type="password"
                                onChange={(event)=>this.changeLoginPassword(event)}
                                fullWidth
                                variant="standard"
                            />
                        </DialogContent>
                        <DialogActions>
                            <Button onClick={()=>this.closeLoginDialog()}>Cancel</Button>
                            <Button onClick={()=>this.clickLogin()}>Login</Button>
                        </DialogActions>
                    </Dialog>                
                    <Dialog open={this.state.signUpOpen} onClose={()=>this.closeSignupDialog()}>
                        <DialogTitle>Sign Up</DialogTitle>
                        <DialogContent>
                            <DialogContentText>
                            <Typography fontSize='1rem'>
                                    {this.state.message}
                                </Typography>                               
                                
                            </DialogContentText>
                            <TextField
                                autoFocus
                                margin="dense"
                                id="name"
                                onChange={(event)=>this.changeSignupUser(event)}
                                label="User Id"
                                type="string"
                                fullWidth
                                variant="standard"
                            />
                            <TextField
                                margin="dense"
                                id="password"
                                label="Password"
                                type="password"
                                onChange={(event)=>this.changeSignupPassword(event)}
                                fullWidth
                                variant="standard"
                            />
                        </DialogContent>
                        <DialogActions>
                            <Button onClick={()=>this.closeSignupDialog()}>Cancel</Button>
                            <Button onClick={()=>this.clickSignup()}>Sign Up</Button>
                        </DialogActions>
                    </Dialog>                

                </div>
            </div>
        );
     }

     userMouseEnter() {
        if (!this.props.loginState.loginId) {
            this.props.updateLoginState({'loginId': this.props.loginState.loginId, 'loginUser': this.props.loginState.loginUser, 'userHover': true});
        }
     }
     userMouseLeave() {
        if (!this.props.loginState.loginId) {
            this.props.updateLoginState({'loginId': this.props.loginState.loginId, 'loginUser': this.props.loginState.loginUser, 'userHover': false});
        }
    }
    signOutUpMouseEnter() {
        this.props.updateLoginState({'loginId': this.props.loginState.loginId, 'loginUser': this.props.loginState.loginUser, 'signUpOutHover': true});
     }
     signOutUpMouseLeave() {
        this.props.updateLoginState({'loginId': this.props.loginState.loginId, 'loginUser': this.props.loginState.loginUser, 'signUpOutHover': false});
    }
    clickPeople() {
        if (!this.props.loginState.loginId) {
            this.showLoginDialog();
        }
    }
    showLoginDialog() {
        this.setState({loginOpen: true, message: "To login, please enter your user name nad password."});
    }
    clickLogin() {
        var url = `./api/login`;
        var body = {
            user_id: this.state.login_user,
            password: this.state.login_password
        }
        RestCall.invoke("POST", url , body, "Unable to load issues.", this.simulateLogin())
        .then (
            (response) => this.loginResponse(response),
            (message) => this.loginErrorResponse(message));
    }

    loginResponse(response) {
        this.setState({loginOpen: false, 'userHover': false});
        this.props.updateLoginState({'loginId': response.user_id, 'loginUser': response.user_id, 'userHover': false});
    }

    loginErrorResponse(response) {
        this.setState({message: response});
    }

    closeLoginDialog() {
        this.setState({loginOpen: false, 'userHover': false});
    }

    clickSignUpOut() {
        if (this.props.loginState.loginId) {
            // Already logged in, so Signout
            this.props.updateLoginState({'loginId': null, 'loginUser': null, 'userHover': false});
        } else {
            // Show sign in dialog
            this.showSignupDialog();
        }
    }

    clickSignup() {
        var url = `./api/signup`;
        var body = {
            user_id: this.state.signupUser,
            password: this.state.signupPassword
        }
        RestCall.invoke("POST", url , body, "Unable to siginup.", this.simulateLogin())
        .then (
            (response) => this.signUpResponse(response),
            (message) => this.signUpErrorResponse(message));
    }

    signUpResponse(response) {
        this.setState({signUpOpen: false, 'userHover': false});
        this.props.updateLoginState({'loginId': response.user_id, 'loginUser': response.user_id, 'userHover': false});
    }

    signUpErrorResponse(response) {
        this.setState({message: response});
    }
    showSignupDialog() {
        this.setState({signUpOpen: true, message: "Enter your desired user and password to sign up."});
    }
    closeSignupDialog() {
        this.setState({signUpOpen: false});
    }
    changeSignupUser(event) {
        this.setState({signupUser: event.target.value})
    }
    changeSignupPassword(event) {
        this.setState({signupPassword: event.target.value})
    }
    changeLoginUser(event) {
        this.setState({login_user: event.target.value})
    }
    changeLoginPassword(event) {
        this.setState({login_password: event.target.value})
    }

    simulateLogin() {
        var result = null;
        if (process.env.NODE_ENV == "development") {
            var jsonString = `
            {
                "user_id": "user_1"
            }
            `;
            result = JSON.parse(jsonString);
        }
        return result;
    }

    simulateSignup() {
        var result = null;
        if (process.env.NODE_ENV == "development") {
            var jsonString = `
            {
                "user_id": "user_1"
            }
            `;
            result = JSON.parse(jsonString);
        }
        return result;
    }

}
 const mapStateToProps = state => {
    var props = {loginState: state.loginState};
    return props;
}
const mapDispatchToProps = (dispatch) => {
   return {
     updateLoginState: (loginState) => dispatch({type: 'updateLoginState', loginState:loginState}),
     setGeneralMessage: (message) => dispatch({type: 'setGeneralMessage', general_message:message})
   }
}
export default connect(mapStateToProps, mapDispatchToProps)(SignInButton);
