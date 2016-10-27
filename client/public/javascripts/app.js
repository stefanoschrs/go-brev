(function(){
	'use strict';

	angular
		.module('BrowserEvent', ['ui.grid'])
		.run(Run)
		.controller('HomeController', HomeController);

	HomeController.$inject = ['$http'];

	function Run() {
		console.log('App Started');
	}

	function HomeController($http){
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
					field: 'error'
				},
				{
					field: 'url'
				},
				{
					field: 'time',
					width: '101'
				}
			]
		};

		$http
			.get('https://0.0.0.0:8080/events')
			.then(function(res){
				return vm.gridOptions.data = res.data.sort((a, b) => -(a.time - b.time));
			})
			.catch(function(err){
				console.error(err);
			});
	}
})();
