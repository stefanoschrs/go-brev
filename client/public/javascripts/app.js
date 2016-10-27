(function(){
	'use strict';

	angular
		.module('BREV', ['ui.grid'])
		.run(Run)
		.constant('SERVER_URL', 'https://0.0.0.0:8080/events')
		.controller('HomeController', HomeController);

	HomeController.$inject = ['$http', 'SERVER_URL'];

	function Run() {
		console.log('BREV Started');
	}

	function HomeController($http, SERVER_URL){
		var vm = this;

		vm.events = [];
		vm.gridOptions = {
			data: vm.events,
			enableHorizontalScrollbar: 0,
			enableVerticalScrollbar: 0,
			columnDefs: [
				{
					field: 'msg'
				},
				{
					field: 'uuid',
					displayName: 'Client',
					width: '310'
				},
				{
					field: 'time',
					width: '105'
				}
			]
		};

		$http
			.get(SERVER_URL)
			.then(function(res){
				return vm.gridOptions.data = res.data.sort((a, b) => -(a.time - b.time));
			})
			.catch(function(err){
				console.error(err);
			});
	}
})();
