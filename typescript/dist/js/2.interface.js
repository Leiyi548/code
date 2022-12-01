class Student {
    constructor(firstName, middleInitial, lastName) {
        this.firstName = firstName;
        this.middleInitial = middleInitial;
        this.lastName = lastName;
        this.fullName = firstName + ' ' + middleInitial + ' ' + lastName;
    }
}
function greeter_interface(person) {
    return 'Hello, ' + person.firstName + ' ' + person.lastName;
}
let user_interface = { firstName: 'Jane', lastName: 'User' };
console.log(user_interface);
console.log(greeter_interface(user_interface));
let red_student = new Student('Jane', 'M.', 'User');
console.log(greeter_interface(red_student));
